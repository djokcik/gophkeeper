package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
)

type LoginPasswordService interface {
	LoadPasswordByLogin(ctx context.Context, username string) (models.LoginPasswordResponseDto, error)
	SaveLoginPassword(ctx context.Context, username string, password string) error
}

type loginPasswordService struct {
	api  RpcService
	user UserService
}

func NewLoginPasswordService(api RpcService, userService UserService) LoginPasswordService {
	return &loginPasswordService{
		api:  api,
		user: userService,
	}
}

func (s loginPasswordService) LoadPasswordByLogin(ctx context.Context, username string) (models.LoginPasswordResponseDto, error) {
	response, err := s.api.LoadPasswordByLogin(ctx, s.user.GetUser(), username)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("LoadPasswordByLogin: invalid load login password")
		return models.LoginPasswordResponseDto{}, err
	}

	return response, err
}

func (s loginPasswordService) SaveLoginPassword(ctx context.Context, username string, password string) error {
	err := s.api.SaveLoginPassword(ctx, s.user.GetUser(), username, password)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveLoginPassword: invalid load login password")
		return err
	}

	return err
}

func (s loginPasswordService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "LoginPasswordService").Logger()

	return &logger
}
