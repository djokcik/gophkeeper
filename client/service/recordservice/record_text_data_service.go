package recordservice

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/service"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
)

type RecordTextDataService interface {
	RemoveRecordByKey(ctx context.Context, key string) error
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordTextData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordTextData) error
}

type recordTextDataService struct {
	api    service.ClientRpcService
	user   service.ClientUserService
	crypto common.CryptoService
}

func NewTextDataService(api service.ClientRpcService, user service.ClientUserService, crypto common.CryptoService) RecordTextDataService {
	return &recordTextDataService{
		api:    api,
		user:   user,
		crypto: crypto,
	}
}

func (s recordTextDataService) RemoveRecordByKey(ctx context.Context, key string) error {
	user := s.user.GetUser()

	err := s.api.RemoveRecordTextDataByKey(ctx, user.Token, key)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("RemoveRecordByKey: invalid load data")
		return err
	}

	return nil
}

func (s recordTextDataService) LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordTextData, error) {
	user := s.user.GetUser()

	encryptedData, err := s.api.LoadRecordTextDataByKey(ctx, user.Token, key)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("LoadRecordByKey: invalid load data")
		return clientmodels.RecordTextData{}, err
	}

	if encryptedData == "" {
		s.Log(ctx).Trace().Msgf("LoadRecordByKey: data by key not found. Key - %s ", key)
		return clientmodels.RecordTextData{}, service.ErrNotFoundLoadData
	}

	var response clientmodels.RecordTextData
	err = s.crypto.DecryptData(ctx, user.Password, encryptedData, &response)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("LoadRecordByKey: invalid decrypt data")
		return clientmodels.RecordTextData{}, err
	}

	s.Log(ctx).Trace().Msg("LoadRecordByKey: success to load")

	return response, err
}

func (s recordTextDataService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordTextData) error {
	user := s.user.GetUser()

	encryptedData, err := s.crypto.EncryptData(ctx, user.Password, data)
	if err != nil {
		return err
	}

	err = s.api.SaveRecordTextData(ctx, user.Token, key, encryptedData)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: invalid SaveRecordTextData")
		return err
	}

	return err
}

func (s recordTextDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordTextDataService").Logger()

	return &logger
}
