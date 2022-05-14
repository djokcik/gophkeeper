package recordservice

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/service"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
)

type RecordBankCardService interface {
	RemoveRecordByKey(ctx context.Context, key string) error
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordBankCardData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordBankCardData) error
}

type recordBankCardService struct {
	api    service.ClientRpcService
	user   service.ClientUserService
	crypto common.CryptoService
}

func NewBankCardService(api service.ClientRpcService, user service.ClientUserService, crypto common.CryptoService) RecordBankCardService {
	return &recordBankCardService{
		api:    api,
		user:   user,
		crypto: crypto,
	}
}

func (s recordBankCardService) RemoveRecordByKey(ctx context.Context, key string) error {
	user := s.user.GetUser()

	err := s.api.RemoveRecordBankCardByKey(ctx, user.Token, key)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("RemoveRecordByKey: invalid load data")
		return err
	}

	return nil
}

func (s recordBankCardService) LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordBankCardData, error) {
	user := s.user.GetUser()

	encryptedData, err := s.api.LoadRecordBankCardByKey(ctx, user.Token, key)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("LoadRecordByKey: invalid load data")
		return clientmodels.RecordBankCardData{}, err
	}

	if encryptedData == "" {
		s.Log(ctx).Trace().Msgf("LoadRecordByKey: data by key not found. Key - %s ", key)
		return clientmodels.RecordBankCardData{}, service.ErrNotFoundLoadData
	}

	var response clientmodels.RecordBankCardData
	err = s.crypto.DecryptData(ctx, user.Password, encryptedData, &response)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("LoadRecordByKey: invalid decrypt data")
		return clientmodels.RecordBankCardData{}, err
	}

	s.Log(ctx).Trace().Msg("LoadRecordByKey: success to load")

	return response, err
}

func (s recordBankCardService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordBankCardData) error {
	user := s.user.GetUser()

	encryptedData, err := s.crypto.EncryptData(ctx, user.Password, data)
	if err != nil {
		return err
	}

	err = s.api.SaveRecordBankCard(ctx, user.Token, key, encryptedData)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: invalid SaveRecordBankCard")
		return err
	}

	return err
}

func (s recordBankCardService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBankCardService").Logger()

	return &logger
}
