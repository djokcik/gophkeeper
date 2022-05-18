package rpchandler

import (
	"context"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/djokcik/gophkeeper/pkg/logging"
)

func (h *RpcHandler) LoadRecordBinaryDataByKeyHandler(dto rpcdto.LoadRecordRequestDto, Reply *string) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call GetUserByToken")
		return err
	}

	*Reply, err = h.recordBinaryData.Load(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call load personal data")
	}

	h.Log(ctx).Trace().Msg("LoadRecordBinaryDataByKeyHandler: load success")

	return nil
}

func (h *RpcHandler) SaveRecordBinaryDataHandler(dto rpcdto.SaveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBinaryDataHandler: err call GetUserByToken")
		return err
	}

	err = h.recordBinaryData.Save(ctx, dto.Key, user.Username, dto.Data)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBinaryDataHandler: err call Save")
		return err
	}

	h.Log(ctx).Trace().Msg("SaveRecordBinaryDataHandler: saved success")

	return nil
}

func (h *RpcHandler) RemoveRecordBinaryDataByKeyHandler(dto rpcdto.RemoveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBinaryDataHandler: err call GetUserByToken")
		return err
	}

	err = h.recordBinaryData.Remove(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBinaryDataHandler: err call Save")
		return err
	}

	h.Log(ctx).Trace().Msg("RemoveRecordBinaryDataByKeyHandler: removed success")

	return nil
}