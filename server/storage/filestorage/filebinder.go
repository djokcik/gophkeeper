package filestorage

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/common"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server/storage"
	"github.com/rs/zerolog"
	"os"
)

var (
	algorithm = map[uint16]string{1: "AES-256-GCM"}
)

const (
	lenVersion             = 2
	lenAlgorithm           = 2
	lenVersionAndAlgorithm = lenAlgorithm + lenAlgorithm
	lenBeforeData          = lenVersionAndAlgorithm
)

//go:generate mockery --name=FileBinder --with-expecter
type FileBinder interface {
	CheckFileExist(ctx context.Context, username string) (bool, error)
	SaveStorage(ctx context.Context, data models.StorageData) error
	ReadStorage(ctx context.Context, username string) (models.StorageData, error)
}

type fileCryptoBinder struct {
	crypto    common.CryptoService
	secretKey string
	baseDir   string
}

func NewFileCryptoBinder(secretKey string, baseDir string) FileBinder {
	crypto := common.NewCryptoService()

	return &fileCryptoBinder{
		crypto:    crypto,
		secretKey: crypto.GenerateHash(secretKey),
		baseDir:   baseDir,
	}
}

func (s fileCryptoBinder) CheckFileExist(ctx context.Context, username string) (bool, error) {
	_, err := os.Stat(s.GetFilename(ctx, username))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (s fileCryptoBinder) ReadStorage(ctx context.Context, username string) (models.StorageData, error) {
	filename := s.GetFilename(ctx, username)

	bytes, err := os.ReadFile(filename)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("ReadStorage: error read file")
		return models.StorageData{}, err
	}

	version := binary.BigEndian.Uint16(bytes[:lenVersion])
	if version != 1 {
		s.Log(ctx).Warn().Err(err).Msg("ReadStorage: invalid version")
		return models.StorageData{}, storage.ErrInvalidFileVersion
	}

	algo := binary.BigEndian.Uint16(bytes[lenVersion:lenVersionAndAlgorithm])
	if _, ok := algorithm[algo]; !ok {
		s.Log(ctx).Warn().Err(err).Msg("ReadStorage: invalid algorithm")
		return models.StorageData{}, storage.ErrNotFoundFileDecrypt
	}

	bytes, err = s.crypto.Decrypt(context.Background(), string(bytes[lenBeforeData:]), s.secretKey)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("ReadStorage: err invalid storage")
		return models.StorageData{}, err
	}

	var data models.StorageData

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("ReadStorage: invalid unmarshal data")
		return models.StorageData{}, err
	}

	return data, nil
}

func (s fileCryptoBinder) SaveStorage(ctx context.Context, data models.StorageData) error {
	filename := s.GetFilename(ctx, data.User.Username)

	marshalData, err := json.Marshal(data)
	if err != nil {
		s.Log(ctx).Trace().Err(err).Msg("SaveStorage: invalid marshal storage data")
		return err
	}

	encryptedData, err := s.crypto.Encrypt(ctx, marshalData, s.secretKey)
	if err != nil {
		s.Log(ctx).Trace().Err(err).Msg("SaveStorage: invalid encrypt data")
		return err
	}

	bytes := make([]byte, lenBeforeData, lenBeforeData+len(encryptedData))
	binary.BigEndian.PutUint16(bytes, 1)              // version
	binary.BigEndian.PutUint16(bytes[lenVersion:], 1) // algorithm
	bytes = append(bytes, []byte(encryptedData)...)

	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveStorage: invalid write to file")
		return err
	}

	return nil
}

func (s fileCryptoBinder) GetFilename(_ context.Context, username string) string {
	return fmt.Sprintf("%s/%s.gkdb", s.baseDir, s.crypto.GenerateHash(username))
}

func (s fileCryptoBinder) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "FileCryptoBinder").Logger()

	return &logger
}
