package recordservice

import (
	"context"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	"github.com/djokcik/gophkeeper/client/service"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=RecordTextDataService --with-expecter
type RecordTextDataService interface {
	RemoveRecordByKey(ctx context.Context, key string) error
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordTextData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordTextData) error
}

type recordTextDataService struct {
	api    service.ClientRpcService
	user   service.ClientUserService
	record ClientRecordService
}

func NewTextDataService(api service.ClientRpcService, user service.ClientUserService, record ClientRecordService) RecordTextDataService {
	return &recordTextDataService{
		api:    api,
		user:   user,
		record: record,
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

	var response clientmodels.RecordTextData
	err := s.record.LoadRecordByKey(ctx, user, &response, func() (string, error) {
		return s.api.LoadRecordTextDataByKey(ctx, user.Token, key)
	})
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("LoadRecordByKey:")
		return clientmodels.RecordTextData{}, err
	}

	return response, nil
}

func (s recordTextDataService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordTextData) error {
	user := s.user.GetUser()

	err := s.record.SaveRecord(ctx, user, data, func(encryptedData string) error {
		return s.api.SaveRecordTextData(ctx, user.Token, key, encryptedData)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s recordTextDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordTextDataService").Logger()

	return &logger
}
