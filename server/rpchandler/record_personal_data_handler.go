package rpchandler

import (
	"context"
	"gophkeeper/models/rpcdto"
)

func (h *RpcHandler) LoadRecordPrivateDataByKeyHandler(dto rpcdto.LoadRecordRequestDto, Reply *string) error {
	ctx := context.Background()

	//authService := h.serviceRegistry.GetAuthService()
	loginPasswordService := h.serviceRegistry.GetLoginPasswordService()

	//if ok, err := authService.Validate(ctx, dto.User); !ok || err != nil {
	//	h.Log(ctx).Error().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: invalid auth")
	//	return err
	//}

	data, err := loginPasswordService.LoadPasswordByLogin(ctx, "test", dto.Key)
	if err != nil {
		h.Log(ctx).Error().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: error LoadRecordPersonalDataByKey")
		return err
	}

	*Reply = data

	h.Log(ctx).Info().Msgf("success load password - %v", data)

	return nil
}

func (h *RpcHandler) SaveRecordPersonalDataHandler(dto rpcdto.SaveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()

	//authService := h.serviceRegistry.GetAuthService()
	loginPasswordService := h.serviceRegistry.GetLoginPasswordService()

	//if ok, err := authService.Validate(ctx, dto.User); !ok || err != nil {
	//	h.Log(ctx).Error().Err(err).Msg("SaveLoginPasswordHandler: invalid auth")
	//	return err
	//}

	err := loginPasswordService.SaveLoginPassword(ctx, "test", dto.Key, dto.Data)
	if err != nil {
		h.Log(ctx).Error().Err(err).Msg("SaveLoginPasswordHandler: error in SaveRecord")
		return err
	}

	h.Log(ctx).Info().Msgf("success save password - %v", dto.Data)
	*Reply = struct{}{}

	return nil
}
