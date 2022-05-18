package common

import (
	"context"
	"errors"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server/storage"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=AuthService --with-expecter

// AuthService provide operations with token
type AuthService interface {
	GetUserByToken(ctx context.Context, token string) (models.User, error)
}

type authService struct {
	secretKey string
	storage   storage.Storage
	authUtils AuthUtilsService
}

func NewAuthService(store storage.Storage, jwtSecretKey string, userUtils AuthUtilsService) AuthService {
	return &authService{
		storage:   store,
		secretKey: jwtSecretKey,
		authUtils: userUtils,
	}
}

// GetUserByToken returns user by token
func (s authService) GetUserByToken(ctx context.Context, token string) (models.User, error) {
	if token == "" {
		s.Log(ctx).Warn().Msg("GetUserByToken: token is empty")
		return models.User{}, ErrUnauthorized
	}

	username, err := s.authUtils.ParseToken(token, s.secretKey)
	if err != nil {
		if errors.Is(err, ErrInvalidAccessToken) {
			s.Log(ctx).Warn().Err(err).Msg("GetUserByToken:")
			return models.User{}, ErrInvalidAccessToken
		}

		s.Log(ctx).Error().Err(err).Msg("GetUserByToken: invalid parse token")
		return models.User{}, err
	}

	return s.storage.UserByUsername(ctx, username)
}

func (s authService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "AuthService").Logger()

	return &logger
}
