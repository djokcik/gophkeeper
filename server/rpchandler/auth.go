package rpchandler

import (
	"context"
	"gophkeeper/models"
	"gophkeeper/models/rpcdto"
)

func (h *RpcHandler) LoginHandler(loginDto rpcdto.LoginDto, Reply *models.GophUser) error {
	ctx := context.Background()

	authService := h.serviceRegistry.GetAuthService()
	user, err := authService.Login(ctx, loginDto.Login, loginDto.Password)
	if err != nil {
		h.Log(ctx).Error().Err(err).Msg("Login: invalid login")
		return err
	}

	*Reply = models.GophUser{
		Username: user.Username,
		Password: user.Password,
	}

	h.Log(ctx).Info().Msgf("success login - %v", Reply)

	return nil
}

func (h *RpcHandler) RegisterHandler(registerDto rpcdto.RegisterDto, Reply *models.GophUser) error {
	ctx := context.Background()

	authService := h.serviceRegistry.GetAuthService()
	user, err := authService.Register(ctx, registerDto.Login, registerDto.Password)
	if err != nil {
		h.Log(ctx).Error().Err(err).Msg("Login: invalid register")
		return err
	}

	*Reply = models.GophUser{
		Username: user.Username,
		Password: user.Password,
	}

	h.Log(ctx).Info().Msgf("success register - %v", Reply)

	return nil
}
