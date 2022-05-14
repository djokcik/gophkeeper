package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/storage"
)

type ServerRecordBinaryDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
	Remove(ctx context.Context, key string, username string) error
}

type recordBinaryDataService struct {
	cfg     server.Config
	storage storage.Storage
	keyLock KeyLockService
}

func NewServerRecordBinaryDataService(cfg server.Config, store storage.Storage, keyLock KeyLockService) ServerRecordBinaryDataService {
	return &recordBinaryDataService{
		cfg:     cfg,
		storage: store,
		keyLock: keyLock,
	}
}

func (s recordBinaryDataService) Save(ctx context.Context, key string, username string, data string) error {
	s.keyLock.Lock(username)
	defer s.keyLock.Unlock(username)

	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Save: invalid read storage")
		return err
	}

	if store.BinaryData == nil {
		store.BinaryData = make(map[string]string)
	}

	store.BinaryData[key] = data

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Save: invalid save data")
		return err
	}

	return nil
}

func (s recordBinaryDataService) Load(ctx context.Context, key string, username string) (string, error) {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Load: invalid read storage")
		return "", err
	}

	return store.BinaryData[key], nil
}

func (s recordBinaryDataService) Remove(ctx context.Context, key string, username string) error {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Remove: invalid read storage")
		return err
	}

	if store.BinaryData == nil {
		return ErrNotFoundRecord
	}

	if _, ok := store.BinaryData[key]; !ok {
		return ErrNotFoundRecord
	}

	delete(store.BinaryData, key)

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Remove: invalid save storage")
		return err
	}

	return nil
}

func (s *recordBinaryDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBinaryDataService").Logger()

	return &logger
}
