package service

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"gophkeeper/models"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/storage"
)

type ServerUserService interface {
	Authenticate(ctx context.Context, login string, password string) (string, error)
	CreateUser(ctx context.Context, username string, password string) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GenerateToken(ctx context.Context, user models.User) (string, error)
}

type userService struct {
	cfg  server.Config
	repo storage.Storage

	auth common.AuthUtilsService
}

func NewAuthService(cfg server.Config, storage storage.Storage, auth common.AuthUtilsService) ServerUserService {
	return &userService{
		cfg:  cfg,
		repo: storage,
		auth: auth,
	}
}

func (u userService) Authenticate(ctx context.Context, login string, password string) (string, error) {
	user, err := u.GetUserByUsername(ctx, login)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			u.Log(ctx).Trace().Err(err).Msg("Authenticate: wrong username")
			return "", ErrWrongPassword
		}

		return "", err
	}

	if err = u.auth.CompareHashAndPassword(password+u.cfg.PasswordPepper, user.Password); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			u.Log(ctx).Trace().Err(err).Msg("Authenticate: wrong password")
			return "", ErrWrongPassword
		}

		return "", err
	}

	token, err := u.GenerateToken(ctx, user)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u userService) CreateUser(ctx context.Context, username string, password string) error {
	user := models.User{Username: username, Password: password}
	err := user.Validate()

	if err != nil {
		u.Log(ctx).Trace().Err(err).Msgf("CreateUser: invalid validate user")
		return err
	}

	user.Password, err = u.auth.HashAndSalt(user.Password, u.cfg.PasswordPepper)
	if err != nil {
		u.Log(ctx).Trace().Err(err).Msgf("error create hash")
		return err
	}

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		u.Log(ctx).Trace().Err(err).Msg("invalid create user")
		return err
	}

	u.Log(ctx).Info().
		Str("Username", user.Username).
		Msg("success created user")

	return nil
}

func (u userService) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	user, err := u.repo.UserByUsername(ctx, username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u userService) GenerateToken(ctx context.Context, user models.User) (string, error) {
	token, err := u.auth.CreateToken(u.cfg.JWTSecretKey, user.Username)
	if err != nil {
		u.Log(ctx).Err(err).Msgf("GenerateToken: error create token")
		return "", err
	}

	return token, nil
}

func (u *userService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerUserService").Logger()

	return &logger
}
