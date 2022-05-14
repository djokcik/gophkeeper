package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/storage"
)

type ServerRecordTextDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
	Remove(ctx context.Context, key string, username string) error
}

type recordTextDataService struct {
	cfg     server.Config
	storage storage.Storage
	keyLock KeyLockService
}

func NewServerRecordTextDataService(cfg server.Config, store storage.Storage, keyLock KeyLockService) ServerRecordTextDataService {
	return &recordTextDataService{
		cfg:     cfg,
		storage: store,
		keyLock: keyLock,
	}
}

func (s recordTextDataService) Save(ctx context.Context, key string, username string, data string) error {
	s.keyLock.Lock(username)
	defer s.keyLock.Unlock(username)

	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Save: invalid read storage")
		return err
	}

	if store.TextData == nil {
		store.TextData = make(map[string]string)
	}

	store.TextData[key] = data

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Save: invalid save data")
		return err
	}

	return nil
}

func (s recordTextDataService) Load(ctx context.Context, key string, username string) (string, error) {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Load: invalid read storage")
		return "", err
	}

	return store.TextData[key], nil
}

func (s recordTextDataService) Remove(ctx context.Context, key string, username string) error {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Remove: invalid read storage")
		return err
	}

	if store.TextData == nil {
		return ErrNotFoundRecord
	}

	if _, ok := store.TextData[key]; !ok {
		return ErrNotFoundRecord
	}

	delete(store.TextData, key)

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Remove: invalid save storage")
		return err
	}

	return nil
}

func (s *recordTextDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordTextDataService").Logger()

	return &logger
}
