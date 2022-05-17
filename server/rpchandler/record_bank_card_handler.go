package rpchandler

import (
	"context"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/djokcik/gophkeeper/pkg/logging"
)

func (h *RpcHandler) LoadRecordBankCardByKeyHandler(dto rpcdto.LoadRecordRequestDto, Reply *string) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	h.Log(ctx).Trace().Msg("LoadRecordBankCardByKeyHandler: start")

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call GetUserByToken")
		return err
	}

	*Reply, err = h.recordBankCard.Load(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("LoadRecordPrivateDataByKeyHandler: err call load personal data")
	}

	h.Log(ctx).Trace().Msg("LoadRecordBankCardByKeyHandler: load success")

	return nil
}

func (h *RpcHandler) SaveRecordBankCardHandler(dto rpcdto.SaveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBankCardHandler: err call GetUserByToken")
		return err
	}

	err = h.recordBankCard.Save(ctx, dto.Key, user.Username, dto.Data)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBankCardHandler: err call Save")
		return err
	}

	h.Log(ctx).Trace().Msg("SaveRecordBankCardHandler: saved success")

	return nil
}

func (h *RpcHandler) RemoveRecordBankCardByKeyHandler(dto rpcdto.RemoveRecordRequestDto, Reply *struct{}) error {
	ctx := context.Background()
	ctx = logging.SetCtxLogger(ctx, h.Log(ctx).With().Logger())

	user, err := h.auth.GetUserByToken(ctx, dto.Token)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBankCardHandler: err call GetUserByToken")
		return err
	}

	err = h.recordBankCard.Remove(ctx, dto.Key, user.Username)
	if err != nil {
		h.Log(ctx).Trace().Err(err).Msg("SaveRecordBankCardHandler: err call Save")
		return err
	}

	h.Log(ctx).Trace().Msg("RemoveRecordBankCardByKeyHandler: removed success")

	return nil
}
