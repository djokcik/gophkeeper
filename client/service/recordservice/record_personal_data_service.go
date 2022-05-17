package recordservice

import (
	"context"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	"github.com/djokcik/gophkeeper/client/service"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=RecordPersonalDataService --with-expecter
type RecordPersonalDataService interface {
	RemoveRecordByKey(ctx context.Context, key string) error
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordPersonalData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordPersonalData) error
}

type recordPersonalDataService struct {
	api    service.ClientRpcService
	user   service.ClientUserService
	record ClientRecordService
}

func NewRecordPersonalDataService(api service.ClientRpcService, user service.ClientUserService, record ClientRecordService) RecordPersonalDataService {
	return &recordPersonalDataService{
		api:    api,
		user:   user,
		record: record,
	}
}

func (s recordPersonalDataService) RemoveRecordByKey(ctx context.Context, key string) error {
	user := s.user.GetUser()

	err := s.api.RemoveRecordPersonalDataByKey(ctx, user.Token, key)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("RemoveRecordByKey: invalid load data")
		return err
	}

	return nil
}

func (s recordPersonalDataService) LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordPersonalData, error) {
	user := s.user.GetUser()

	var response clientmodels.RecordPersonalData
	err := s.record.LoadRecordByKey(ctx, user, &response, func() (string, error) {
		return s.api.LoadRecordPersonalDataByKey(ctx, user.Token, key)
	})
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("LoadRecordByKey:")
		return clientmodels.RecordPersonalData{}, err
	}

	return response, nil
}

func (s recordPersonalDataService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordPersonalData) error {
	user := s.user.GetUser()

	err := s.record.SaveRecord(ctx, user, data, func(encryptedData string) error {
		return s.api.SaveRecordPersonalData(ctx, user.Token, key, encryptedData)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s recordPersonalDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordPersonalDataService").Logger()

	return &logger
}
