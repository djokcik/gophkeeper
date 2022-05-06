package filestorage

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/storage"
)

func NewFileStorage(cfg server.Config) storage.Storage {
	return &fileStorage{
		bind: NewFileCryptoBinder(cfg.DBSecretKey, cfg.StorePath),
	}
}

type fileStorage struct {
	bind FileBinder
}

func (s fileStorage) Read(ctx context.Context, username string) (models.StorageData, error) {
	return s.bind.ReadStorage(ctx, username)
}

func (s fileStorage) Save(ctx context.Context, data models.StorageData) error {
	return s.bind.SaveStorage(ctx, data)
}

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
