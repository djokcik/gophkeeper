package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/storage"
)

type ServerRecordBankCardDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
	Remove(ctx context.Context, key string, username string) error
}

type recordBankCardDataService struct {
	cfg     server.Config
	storage storage.Storage
	keyLock KeyLockService
}

func NewServerRecordBankCardDataService(cfg server.Config, store storage.Storage, keyLock KeyLockService) ServerRecordBankCardDataService {
	return &recordBankCardDataService{
		cfg:     cfg,
		storage: store,
		keyLock: keyLock,
	}
}

func (s recordBankCardDataService) Save(ctx context.Context, key string, username string, data string) error {
	s.keyLock.Lock(username)
	defer s.keyLock.Unlock(username)

	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Save: invalid read storage")
		return err
	}

	if store.BankCardData == nil {
		store.BankCardData = make(map[string]string)
	}

	store.BankCardData[key] = data

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Save: invalid save data")
		return err
	}

	return nil
}

func (s recordBankCardDataService) Load(ctx context.Context, key string, username string) (string, error) {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Load: invalid read storage")
		return "", err
	}

	return store.BankCardData[key], nil
}

func (s recordBankCardDataService) Remove(ctx context.Context, key string, username string) error {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Remove: invalid read storage")
		return err
	}

	if store.BankCardData == nil {
		return ErrNotFoundRecord
	}

	if _, ok := store.BankCardData[key]; !ok {
		return ErrNotFoundRecord
	}

	delete(store.BankCardData, key)

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Remove: invalid save storage")
		return err
	}

	return nil
}

func (s *recordBankCardDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBankCardDataService").Logger()

	return &logger
}
