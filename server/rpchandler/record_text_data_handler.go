package rpchandler

import (
	"context"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/djokcik/gophkeeper/pkg/logging"
)

// LoadRecordTextDataByKeyHandler handler for load record text data
func (h *RPCHandler) LoadRecordTextDataByKeyHandler(dto rpcdto.LoadRecordRequestDto, Reply *string) error {
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

	h.Log(ctx).Trace().Msg("LoadRecordTextDataByKeyHandler: load success")

	return nil
}

// SaveRecordTextDataHandler handler for save record text data
func (h *RPCHandler) SaveRecordTextDataHandler(dto rpcdto.SaveRecordRequestDto, Reply *struct{}) error {
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

	h.Log(ctx).Trace().Msg("SaveRecordTextDataHandler: saved success")

	return nil
}

// RemoveRecordTextDataByKeyHandler handler for remove record text data
func (h *RPCHandler) RemoveRecordTextDataByKeyHandler(dto rpcdto.RemoveRecordRequestDto, Reply *struct{}) error {
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

	h.Log(ctx).Trace().Msg("RemoveRecordTextDataByKeyHandler: removed success")

	return nil
}
