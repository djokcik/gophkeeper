package rpchandler

import (
	"context"
	"gophkeeper/models/rpcdto"
	"gophkeeper/pkg/logging"
)

func (h *RpcHandler) LoadRecordTextDataByKeyHandler(dto rpcdto.LoadRecordRequestDto, Reply *string) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call GetUserByToken")
		return err
	}

	*Reply, err = h.recordTextData.Load(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call load personal data")
	}

	return nil
}

func (h *RpcHandler) SaveRecordTextDataHandler(dto rpcdto.SaveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordTextDataHandler: err call GetUserByToken")
		return err
	}

	err = h.recordTextData.Save(ctx, dto.Key, user.Username, dto.Data)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordTextDataHandler: err call Save")
		return err
	}

	return nil
}

func (h *RpcHandler) RemoveRecordTextDataByKeyHandler(dto rpcdto.RemoveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordTextDataHandler: err call GetUserByToken")
		return err
	}

	err = h.recordTextData.Remove(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordTextDataHandler: err call Save")
		return err
	}

	return nil
}
