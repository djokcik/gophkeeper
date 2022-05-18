package service

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server/storage"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=ServerRecordService --with-expecter

// ServerRecordService provides base methods for all types record operations
type ServerRecordService interface {
	Save(ctx context.Context, username string, updateFn func(store *models.StorageData) error) error
	Load(ctx context.Context, username string, loadFn func(store models.StorageData) string) (string, error)
	Remove(ctx context.Context, username string, removeFn func(store *models.StorageData) error) error
}

type recordService struct {
	keyLock KeyLockService
	storage storage.Storage
}

func NewRecordService(keyLock KeyLockService, storage storage.Storage) ServerRecordService {
	return &recordService{
		keyLock: keyLock,
		storage: storage,
	}
}

// Save is base save record
func (s recordService) Save(ctx context.Context, username string, updateFn func(store *models.StorageData) error) error {
	s.keyLock.Lock(username)
	defer s.keyLock.Unlock(username)

	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Save: invalid read storage")
		return err
	}

	err = updateFn(&store)
	if err != nil {
		s.Log(ctx).Error().Msg("Save: failed to update store")
		return err
	}

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Save: invalid save data")
		return err
	}

	return nil
}

// Load is base load record
func (s recordService) Load(ctx context.Context, username string, loadFn func(store models.StorageData) string) (string, error) {
	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Load: invalid read storage")
		return "", err
	}

	return loadFn(store), nil
}

// Remove is base remove record
func (s recordService) Remove(ctx context.Context, username string, removeFn func(store *models.StorageData) error) error {
	s.keyLock.Lock(username)
	defer s.keyLock.Unlock(username)

	store, err := s.storage.Read(ctx, username)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("Remove: invalid read storage")
		return err
	}

	err = removeFn(&store)
	if err != nil {
		return err
	}

	err = s.storage.Save(ctx, store)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Remove: invalid save storage")
		return err
	}

	return nil
}

func (s *recordService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordService").Logger()

	return &logger
}
