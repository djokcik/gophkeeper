package rpchandler

import (
	"context"
	"gophkeeper/models"
	"gophkeeper/models/rpcdto"
)

func (h *RpcHandler) LoadPasswordByLoginHandler(dto rpcdto.LoadRecordRequestDto, Reply *models.LoginPasswordResponseDto) error {
	ctx := context.Background()

	authService := h.serviceRegistry.GetAuthService()
	loginPasswordService := h.serviceRegistry.GetLoginPasswordService()

	if ok, err := authService.Validate(ctx, dto.User); !ok || err != nil {
		h.Log(ctx).Error().Err(err).Msg("LoadPasswordByLoginHandler: invalid auth")
		return err
	}

	data, err := loginPasswordService.LoadPasswordByLogin(ctx, dto.User, dto.Login)
	if err != nil {
		h.Log(ctx).Error().Err(err).Msg("LoadPasswordByLoginHandler: error LoadRecordPersonalDataByKey")
		return err
	}

	*Reply = data

	h.Log(ctx).Info().Msgf("success load password - %v", Reply)

	return nil
}

func (h *RpcHandler) SaveLoginPasswordHandler(dto rpcdto.SaveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()

	authService := h.serviceRegistry.GetAuthService()
	loginPasswordService := h.serviceRegistry.GetLoginPasswordService()

	if ok, err := authService.Validate(ctx, dto.User); !ok || err != nil {
		h.Log(ctx).Error().Err(err).Msg("SaveLoginPasswordHandler: invalid auth")
		return err
	}

	err := loginPasswordService.SaveLoginPassword(ctx, dto.User, dto.Login, dto.Password)
	if err != nil {
		h.Log(ctx).Error().Err(err).Msg("SaveLoginPasswordHandler: error in SaveRecord")
		return err
	}

	h.Log(ctx).Info().Msgf("success save password - %v", dto.Login)

	return nil
}
