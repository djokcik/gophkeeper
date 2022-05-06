package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/storage"
)

type ServerRecordPersonalDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
}

type recordPersonalDataService struct {
	cfg     server.Config
	storage storage.Storage
	keyLock KeyLockService
}

func NewServerRecordPersonalDataService(cfg server.Config, store storage.Storage, keyLock KeyLockService) ServerRecordPersonalDataService {
	return &recordPersonalDataService{
		cfg:     cfg,
		storage: store,
		keyLock: keyLock,
	}
}

func (s recordPersonalDataService) Save(ctx context.Context, key string, username string, data string) error {
	s.keyLock.Lock(username)
	defer s.keyLock.Unlock(username)

	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Save: invalid read storage")
		return err
	}

	if store.PersonalData == nil {
		store.PersonalData = make(map[string]string)
	}

	store.PersonalData[key] = data

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Save: invalid save data")
		return err
	}

	return nil
}

func (s recordPersonalDataService) Load(ctx context.Context, key string, username string) (string, error) {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Load: invalid read storage")
		return "", err
	}

	return store.PersonalData[key], nil
}

func (s *recordPersonalDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordPersonalDataService").Logger()

	return &logger
}
