package filestorage

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server"
	"github.com/djokcik/gophkeeper/server/storage"
	"github.com/rs/zerolog"
)

func NewFileStorage(cfg server.Config) storage.Storage {
	return &fileStorage{
		bind: NewFileCryptoBinder(cfg.DBSecretKey, cfg.StorePath),
	}
}

// fileStorage is implements Storage
type fileStorage struct {
	bind FileBinder
}

// Read returns storageData by username
func (s fileStorage) Read(ctx context.Context, username string) (models.StorageData, error) {
	return s.bind.ReadStorage(ctx, username)
}

// Save is saved storageData in DB
func (s fileStorage) Save(ctx context.Context, data models.StorageData) error {
	return s.bind.SaveStorage(ctx, data)
}

// UserByUsername returns user by username
func (s fileStorage) UserByUsername(ctx context.Context, username string) (models.User, error) {
	fileExists, err := s.bind.CheckFileExist(ctx, username)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("UserByUsername: invalid call CheckFileExist")
		return models.User{}, err
	}

	if !fileExists {
		s.Log(ctx).Warn().Err(err).Msg("UserByUsername: user not found")
		return models.User{}, storage.ErrNotFound
	}

	data, err := s.bind.ReadStorage(ctx, username)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("UserByUsername: err in ReadStorage")
		return models.User{}, err
	}

	return data.User, nil
}

// CreateUser is created user in DB
func (s fileStorage) CreateUser(ctx context.Context, user models.User) error {
	storageData := models.StorageData{User: user}

	fileExists, err := s.bind.CheckFileExist(ctx, storageData.User.Username)
	if err != nil {
		s.Log(ctx).Warn().Msg("CreateUser: err in CheckFileExist")
		return err
	}

	if fileExists {
		s.Log(ctx).Error().Err(err).Msg("CreateUser: login already exists")
		return storage.ErrLoginAlreadyExists
	}

	err = s.bind.SaveStorage(ctx, storageData)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("CreateUser: invalid save storage data")
		return err
	}

	return nil
}

func (s fileStorage) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "FileStorage").Logger()

	return &logger
}
