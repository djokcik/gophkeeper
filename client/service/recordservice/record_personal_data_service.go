package recordservice

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/service"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
)

type RecordPersonalDataService interface {
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordPersonalData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordPersonalData) error
}

type loginPasswordService struct {
	api    service.ClientRpcService
	user   service.ClientUserService
	crypto common.CryptoService
}

func NewLoginPasswordService(api service.ClientRpcService, user service.ClientUserService, crypto common.CryptoService) RecordPersonalDataService {
	return &loginPasswordService{
		api:    api,
		user:   user,
		crypto: crypto,
	}
}

func (s loginPasswordService) LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordPersonalData, error) {
	user := s.user.GetUser()

	encryptedData, err := s.api.LoadRecordPrivateDataByKey(ctx, user.Token, key)
	if err != nil {
		return clientmodels.RecordPersonalData{}, err
	}

	var response clientmodels.RecordPersonalData
	err = s.crypto.DecryptData(ctx, user, encryptedData, &response)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("LoadRecordByKey: invalid decrypt data")
		return clientmodels.RecordPersonalData{}, err
	}

	s.Log(ctx).Trace().Msg("LoadRecordByKey: success to load")

	return response, err
}

func (s loginPasswordService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordPersonalData) error {
	user := s.user.GetUser()

	encryptedData, err := s.crypto.EncryptData(ctx, user, data)
	if err != nil {
		return err
	}

	err = s.api.SaveRecordPersonalData(ctx, user.Token, key, encryptedData)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: invalid SaveRecordPersonalData")
		return err
	}

	return err
}

func (s loginPasswordService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "RecordPersonalDataService").Logger()

	return &logger
}
