package rpchandler

import (
	"context"
	"gophkeeper/models/rpcdto"
	"gophkeeper/pkg/logging"
)

func (h *RpcHandler) LoadRecordPersonalDataByKeyHandler(dto rpcdto.LoadRecordRequestDto, Reply *string) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call GetUserByToken")
		return err
	}

	*Reply, err = h.recordPersonalData.Load(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call load personal data")
	}

	return nil
}

func (h *RpcHandler) SaveRecordPersonalDataHandler(dto rpcdto.SaveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordPersonalDataHandler: err call GetUserByToken")
		return err
	}

	err = h.recordPersonalData.Save(ctx, dto.Key, user.Username, dto.Data)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordPersonalDataHandler: err call Save")
		return err
	}

	return nil
}

func (h *RpcHandler) RemoveRecordPersonalDataByKeyHandler(dto rpcdto.RemoveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordPersonalDataHandler: err call GetUserByToken")
		return err
	}

	err = h.recordPersonalData.Remove(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordPersonalDataHandler: err call Save")
		return err
	}

	return nil
}
