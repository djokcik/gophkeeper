package service

import (
	"context"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
)

type AuthService interface {
	Login(ctx context.Context, username string, password string) (models.GophUser, error)
	Register(ctx context.Context, username string, password string) (models.GophUser, error)
}

// Ensure service implements interface
var _ AuthService = (*authService)(nil)

type authService struct {
	api RpcService
}

func NewAuthService(api RpcService) AuthService {
	return &authService{
		api: api,
	}
}

func (s authService) Login(ctx context.Context, username string, password string) (models.GophUser, error) {
	user, err := s.api.Login(ctx, username, password)
	if err != nil {
		logging.NewFileLogger().Warn().Err(err).Msg("invalid login")
	}

	return user, err
}

func (s authService) Register(ctx context.Context, username string, password string) (models.GophUser, error) {
	user, err := s.api.Register(ctx, username, password)
	if err != nil {
		logging.NewFileLogger().Warn().Err(err).Msg("invalid login")
	}

	return user, err
}
