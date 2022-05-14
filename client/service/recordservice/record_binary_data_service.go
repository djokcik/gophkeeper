package recordservice

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/service"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
)

type RecordBinaryDataService interface {
	RemoveRecordByKey(ctx context.Context, key string) error
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordBinaryData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordBinaryData) error
}

type recordBinaryDataService struct {
	api    service.ClientRpcService
	user   service.ClientUserService
	crypto common.CryptoService
}

func NewBinaryDataService(api service.ClientRpcService, user service.ClientUserService, crypto common.CryptoService) RecordBinaryDataService {
	return &recordBinaryDataService{
		api:    api,
		user:   user,
		crypto: crypto,
	}
}

func (s recordBinaryDataService) RemoveRecordByKey(ctx context.Context, key string) error {
	user := s.user.GetUser()

	err := s.api.RemoveRecordBinaryDataByKey(ctx, user.Token, key)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("RemoveRecordByKey: invalid load data")
		return err
	}

	return nil
}

func (s recordBinaryDataService) LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordBinaryData, error) {
	user := s.user.GetUser()

	encryptedData, err := s.api.LoadRecordBinaryDataByKey(ctx, user.Token, key)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("LoadRecordByKey: invalid load data")
		return clientmodels.RecordBinaryData{}, err
	}

	if encryptedData == "" {
		s.Log(ctx).Trace().Msgf("LoadRecordByKey: data by key not found. Key - %s ", key)
		return clientmodels.RecordBinaryData{}, service.ErrNotFoundLoadData
	}

	var response clientmodels.RecordBinaryData
	err = s.crypto.DecryptData(ctx, user.Password, encryptedData, &response)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("LoadRecordByKey: invalid decrypt data")
		return clientmodels.RecordBinaryData{}, err
	}

	s.Log(ctx).Trace().Msg("LoadRecordByKey: success to load")

	return response, err
}

func (s recordBinaryDataService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordBinaryData) error {
	user := s.user.GetUser()

	encryptedData, err := s.crypto.EncryptData(ctx, user.Password, data)
	if err != nil {
		return err
	}

	err = s.api.SaveRecordBinaryData(ctx, user.Token, key, encryptedData)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: invalid SaveRecordBinaryData")
		return err
	}

	return err
}

func (s recordBinaryDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBinaryDataService").Logger()

	return &logger
}
