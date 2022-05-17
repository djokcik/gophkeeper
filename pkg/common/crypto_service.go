package common

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/logging"
	"io"
)

//go:generate mockery --name=CryptoService --with-expecter
type CryptoService interface {
	GenerateHash(value string) string
	Encrypt(ctx context.Context, data []byte, key string) (string, error)
	Decrypt(ctx context.Context, encryptedString string, key string) ([]byte, error)

	EncryptData(ctx context.Context, userPassword string, data interface{}) (string, error)
	DecryptData(ctx context.Context, userPassword string, encryptedData string, response interface{}) error
}

type cryptoService struct {
}

func NewCryptoService() CryptoService {
	return &cryptoService{}
}

func (s cryptoService) GenerateHash(key string) string {
	h := sha256.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

func (s cryptoService) Encrypt(ctx context.Context, data []byte, keyString string) (string, error) {
	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := data

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Encrypt: invalid NewCipher")
		return "", err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Encrypt: invalid NewGCM")
		return "", err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func (s cryptoService) Decrypt(ctx context.Context, encryptedString string, keyString string) ([]byte, error) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Decrypt: invalid NewCipher")
		return []byte{}, err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Decrypt: invalid NewGCM")
		return []byte{}, err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Decrypt: invalid NewGCM")
		return []byte{}, err
	}

	return plaintext, nil
}

func (s cryptoService) EncryptData(ctx context.Context, userPassword string, data interface{}) (string, error) {
	key := s.GenerateHash(userPassword)

	bytes, err := json.Marshal(data)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("EncryptData: invalid Marshal")
		return "", err
	}

	encryptedData, err := s.Encrypt(ctx, bytes, key)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("EncryptData: invalid Encrypt")
		return "", err
	}

	return encryptedData, nil
}

func (s cryptoService) DecryptData(ctx context.Context, userPassword string, encryptedData string, response interface{}) error {
	key := s.GenerateHash(userPassword)

	bytes, err := s.Decrypt(ctx, encryptedData, key)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("DecryptData: invalid decrypt data")
		return err
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("DecryptData: invalid unmarshal decrypted data")
		return err
	}

	return nil
}

func (s cryptoService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "CryptoService").Logger()

	return &logger
}
