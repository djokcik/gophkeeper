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
	SaveLoginPassword(ctx context.Context, username string, key string, data string) error
	LoadPasswordByLogin(ctx context.Context, username string, key string) (string, error)
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

func (s loginPasswordService) SaveLoginPassword(ctx context.Context, username string, key string, dataStr string) error {
	filename := fmt.Sprintf("%s/%s.txt", s.Cfg.StorePath, username)

	if username == "" || key == "" {
		return errors.New("invalid validate: username или password неверный")
	}

	_, err := os.Stat(filename)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: error in os.Stat")
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

	data.LoginPasswordMap[key] = dataStr
	err = s.FileStorage.Write(ctx, filename, data)
	if err != nil {
		return err
	}

	return nil
}

func (s loginPasswordService) LoadPasswordByLogin(ctx context.Context, username string, key string) (string, error) {
	filename := fmt.Sprintf("%s/%s.txt", s.Cfg.StorePath, username)

	_, err := os.Stat(filename)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("LoadPasswordByLogin: error in os.Stat")
		return "", err
	}

	var data models.StorageData

	err = s.FileStorage.Read(ctx, filename, &data)
	if err != nil {
		return "", err
	}

	value, ok := data.LoginPasswordMap[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("не найден пароль для %s", username))
	}

	return value, nil
}

func (s *loginPasswordService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "LoginPasswordService").Logger()

	return &logger
}
