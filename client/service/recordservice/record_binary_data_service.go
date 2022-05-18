package recordservice

import (
	"context"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	"github.com/djokcik/gophkeeper/client/service"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=RecordBinaryDataService --with-expecter
type RecordBinaryDataService interface {
	RemoveRecordByKey(ctx context.Context, key string) error
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordBinaryData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordBinaryData) error
}

type recordBinaryDataService struct {
	api    service.ClientRPCService
	user   service.ClientUserService
	record ClientRecordService
}

func NewBinaryDataService(api service.ClientRPCService, user service.ClientUserService, record ClientRecordService) RecordBinaryDataService {
	return &recordBinaryDataService{
		api:    api,
		user:   user,
		record: record,
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

	var response clientmodels.RecordBinaryData
	err := s.record.LoadRecordByKey(ctx, user, &response, func() (string, error) {
		return s.api.LoadRecordBinaryDataByKey(ctx, user.Token, key)
	})
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("LoadRecordByKey:")
		return clientmodels.RecordBinaryData{}, err
	}

	return response, nil
}

func (s recordBinaryDataService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordBinaryData) error {
	user := s.user.GetUser()

	err := s.record.SaveRecord(ctx, user, data, func(encryptedData string) error {
		return s.api.SaveRecordBinaryData(ctx, user.Token, key, encryptedData)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s recordBinaryDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBinaryDataService").Logger()

	return &logger
}
