package rpchandler

import (
	"context"
	"errors"
	"gophkeeper/models"
	"gophkeeper/models/rpcdto"
)

func (h *RpcHandler) Login(loginDto rpcdto.LoginDto, Reply *models.GophUser) error {
	ctx := context.Background()

	if loginDto.Login == "test" {
		h.Log(ctx).Warn().Msg("invalid login")
		return errors.New("invalid login")
	}

	if loginDto.Password == "test" {
		return errors.New("invalid password")
	}

	*Reply = models.GophUser{
		Username: loginDto.Login,
		Password: loginDto.Password,
	}

	h.Log(ctx).Info().Msgf("success login - %v", Reply)

	return nil
}

func (h *RpcHandler) Register(registerDto rpcdto.RegisterDto, Reply *models.GophUser) error {
	ctx := context.Background()

	if registerDto.Login == "test" {
		h.Log(ctx).Warn().Msg("invalid login")
		return errors.New("invalid login")
	}

	if registerDto.Password == "test" {
		return errors.New("invalid password")
	}

	*Reply = models.GophUser{
		Username: registerDto.Login,
		Password: registerDto.Password,
	}

	h.Log(ctx).Info().Msgf("success register - %v", Reply)

	return nil
}
