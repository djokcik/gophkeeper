package service

import (
	"context"
	"errors"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=ClientAuthService

// ClientAuthService provides methods for authorise
type ClientAuthService interface {
	SignIn(ctx context.Context, username string, password string) (models.ClientUser, error)
	Register(ctx context.Context, username string, password string) (models.ClientUser, error)
}

// Ensure service implements interface
var _ ClientAuthService = (*authService)(nil)

type authService struct {
	api ClientRPCService
}

func NewClientAuthService(api ClientRPCService) ClientAuthService {
	return &authService{
		api: api,
	}
}

// SignIn user using username and password
func (s authService) SignIn(ctx context.Context, username string, password string) (models.ClientUser, error) {
	token, err := s.api.Login(ctx, username, password)
	if err != nil && !errors.Is(err, ErrAnonymousUser) {
		s.Log(ctx).Warn().Err(err).Msg("SignIn:")
		return models.ClientUser{}, err
	}

	return models.ClientUser{Username: username, Password: password, Token: token}, nil
}

// Register user using username and password
func (s authService) Register(ctx context.Context, username string, password string) (models.ClientUser, error) {
	token, err := s.api.Register(ctx, username, password)
	if err != nil && !errors.Is(err, ErrAnonymousUser) {
		s.Log(ctx).Warn().Err(err).Msg("Register:")
		return models.ClientUser{}, err
	}

	return models.ClientUser{Username: username, Password: password, Token: token}, nil
}

func (s authService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ClientAuthService").Logger()

	return &logger
}
