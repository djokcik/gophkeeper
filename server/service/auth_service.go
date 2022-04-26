package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/storage"
	"os"
)

type AuthService interface {
	Validate(ctx context.Context, user models.GophUser) (bool, error)
	Login(ctx context.Context, username string, password string) (models.GophUser, error)
	Register(ctx context.Context, username string, password string) (models.GophUser, error)
}

type authService struct {
	Cfg         server.Config
	FileStorage storage.FileStorage
}

func NewAuthService(cfg server.Config, fileStorage storage.FileStorage) AuthService {
	return &authService{
		Cfg:         cfg,
		FileStorage: fileStorage,
	}
}

func (s authService) Validate(ctx context.Context, user models.GophUser) (bool, error) {
	filename := fmt.Sprintf("%s/%s.txt", s.Cfg.StorePath, user.Username)

	_, err := os.Stat(filename)
	if err != nil {
		return false, err
	}

	var data models.StorageData

	err = s.FileStorage.Read(ctx, filename, &data)
	if err != nil {
		return false, err
	}

	return data.User.Password == user.Password, nil
}

func (s authService) Login(ctx context.Context, username string, password string) (models.GophUser, error) {
	filename := fmt.Sprintf("%s/%s.txt", s.Cfg.StorePath, username)

	_, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return models.GophUser{}, errors.New("неверный логин")
		}

		s.Log(ctx).Error().Err(err).Msg("Login: invalid stat file")
		return models.GophUser{}, err
	}

	var data models.StorageData

	err = s.FileStorage.Read(ctx, filename, &data)
	if err != nil {
		return models.GophUser{}, err
	}

	if data.User.Password != password {
		return models.GophUser{}, errors.New("неверный пароль")
	}

	return data.User, nil
}

func (s authService) Register(ctx context.Context, username string, password string) (models.GophUser, error) {
	filename := fmt.Sprintf("%s/%s.txt", s.Cfg.StorePath, username)

	_, err := os.Stat(filename)
	if err == nil || !errors.Is(err, os.ErrNotExist) {
		if !errors.Is(err, os.ErrNotExist) {
			return models.GophUser{}, errors.New("такой пользователь уже зарегистрирован")
		}

		s.Log(ctx).Error().Err(err).Msg("Register: invalid stat file")
		return models.GophUser{}, err
	}

	data := models.StorageData{User: models.GophUser{
		Username: username,
		Password: password,
	}}

	err = s.FileStorage.Write(ctx, filename, data)
	if err != nil {
		return models.GophUser{}, err
	}

	return data.User, nil
}

func (s *authService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "AuthService").Logger()

	return &logger
}
