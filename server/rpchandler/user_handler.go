package rpchandler

import (
	"context"
	"errors"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server/service"
	"github.com/djokcik/gophkeeper/server/storage"
)

func (h *RpcHandler) SignInHandler(loginDto rpcdto.LoginDto, Reply *string) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	token, err := h.user.Authenticate(ctx, loginDto.Login, loginDto.Password)
	if err != nil {
		if errors.Is(err, service.ErrWrongPassword) {
			h.Log(ctx).Trace().Err(err).Msg("SignInHandler: invalid password")
			return err
		}

		h.Log(ctx).Warn().Err(err).Msg("SignInHandler: authenticate error")
		return err
	}

	*Reply = token

	h.Log(ctx).Info().Msgf("RegisterHandler: success login - %v", loginDto.Login)

	return nil
}

func (h *RpcHandler) RegisterHandler(registerDto rpcdto.RegisterDto, Reply *string) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	err := h.user.CreateUser(ctx, registerDto.Login, registerDto.Password)
	if err != nil {
		if errors.Is(err, storage.ErrLoginAlreadyExists) {
			h.Log(ctx).Trace().Err(err).Msg("RegisterHandler: login already exists")
			return err
		}

		h.Log(ctx).Trace().Err(err).Msg("RegisterHandler: failed created user")
		return err
	}

	user, err := h.user.GetUserByUsername(ctx, registerDto.Login)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("RegisterHandler: failed to find by username")
		return err
	}

	token, err := h.user.GenerateToken(ctx, user)
	if err != nil {
		h.Log(ctx).Error().Err(err).Msg("RegisterHandler: invalid generate token")
		return err
	}

	*Reply = token

	h.Log(ctx).Info().Msgf("RegisterHandler: success register - %v", registerDto.Login)

	return nil
}
