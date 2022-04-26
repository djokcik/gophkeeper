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

type LoginPasswordService interface {
	SaveLoginPassword(ctx context.Context, user models.GophUser, username string, password string) error
	LoadPasswordByLogin(ctx context.Context, user models.GophUser, username string) (models.LoginPasswordResponseDto, error)
}

type loginPasswordService struct {
	Cfg         server.Config
	FileStorage storage.FileStorage
}

func NewLoginPasswordService(cfg server.Config, fileStorage storage.FileStorage) LoginPasswordService {
	return &loginPasswordService{
		Cfg:         cfg,
		FileStorage: fileStorage,
	}
}

func (s loginPasswordService) SaveLoginPassword(ctx context.Context, user models.GophUser, username string, password string) error {
	filename := fmt.Sprintf("%s/%s.txt", s.Cfg.StorePath, user.Username)

	if username == "" || password == "" {
		return errors.New("invalid validate: username или password неверный")
	}

	_, err := os.Stat(filename)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveLoginPassword: error in os.Stat")
		return err
	}

	var data models.StorageData

	err = s.FileStorage.Read(ctx, filename, &data)
	if err != nil {
		return err
	}

	if data.LoginPasswordMap == nil {
		data.LoginPasswordMap = make(map[string]string)
	}

	data.LoginPasswordMap[username] = password
	err = s.FileStorage.Write(ctx, filename, data)
	if err != nil {
		return err
	}

	return nil
}

func (s loginPasswordService) LoadPasswordByLogin(ctx context.Context, user models.GophUser, username string) (models.LoginPasswordResponseDto, error) {
	filename := fmt.Sprintf("%s/%s.txt", s.Cfg.StorePath, user.Username)

	_, err := os.Stat(filename)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveLoginPassword: error in os.Stat")
		return models.LoginPasswordResponseDto{}, err
	}

	var data models.StorageData

	err = s.FileStorage.Read(ctx, filename, &data)
	if err != nil {
		return models.LoginPasswordResponseDto{}, err
	}

	password, ok := data.LoginPasswordMap[username]
	if !ok {
		return models.LoginPasswordResponseDto{}, errors.New(fmt.Sprintf("не найден пароль для %s", username))
	}

	return models.LoginPasswordResponseDto{Username: username, Password: password}, nil
}

func (s *loginPasswordService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "LoginPasswordService").Logger()

	return &logger
}
