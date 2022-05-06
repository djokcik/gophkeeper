package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
)

type ClientAuthService interface {
	SignIn(ctx context.Context, username string, password string) (models.ClientUser, error)
	Register(ctx context.Context, username string, password string) (models.ClientUser, error)
}

// Ensure service implements interface
var _ ClientAuthService = (*authService)(nil)

type authService struct {
	api    ClientRpcService
	crypto common.CryptoService
}

func NewClientAuthService(api ClientRpcService, crypto common.CryptoService) ClientAuthService {
	return &authService{
		api:    api,
		crypto: crypto,
	}
}

func (s authService) SignIn(ctx context.Context, username string, password string) (models.ClientUser, error) {
	token, err := s.api.Login(ctx, username, password)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("SignIn:")
		return models.ClientUser{}, err
	}

	return models.ClientUser{Username: username, Password: password, Token: token}, err
}

func (s authService) Register(ctx context.Context, username string, password string) (models.ClientUser, error) {
	token, err := s.api.Register(ctx, username, password)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("CreateUser:")
		return models.ClientUser{}, err
	}

	return models.ClientUser{Username: username, Password: password, Token: token}, err
}

func (s authService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ClientAuthService").Logger()

	return &logger
}
