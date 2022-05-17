package recordservice

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/service"
	"gophkeeper/pkg/logging"
)

//go:generate mockery --name=RecordBankCardService --with-expecter

type RecordBankCardService interface {
	RemoveRecordByKey(ctx context.Context, key string) error
	LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordBankCardData, error)
	SaveRecord(ctx context.Context, key string, data clientmodels.RecordBankCardData) error
}

type recordBankCardService struct {
	api    service.ClientRpcService
	user   service.ClientUserService
	record ClientRecordService
}

func NewBankCardService(api service.ClientRpcService, user service.ClientUserService, record ClientRecordService) RecordBankCardService {
	return &recordBankCardService{
		api:    api,
		user:   user,
		record: record,
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

	var response clientmodels.RecordBankCardData
	err := s.record.LoadRecordByKey(ctx, user, &response, func() (string, error) {
		return s.api.LoadRecordBankCardByKey(ctx, user.Token, key)
	})
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("LoadRecordByKey:")
		return clientmodels.RecordBankCardData{}, err
	}

	return response, nil
}

func (s recordBankCardService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordBankCardData) error {
	user := s.user.GetUser()

	err := s.record.SaveRecord(ctx, user, data, func(encryptedData string) error {
		return s.api.SaveRecordBankCard(ctx, user.Token, key, encryptedData)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s recordBankCardService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBankCardService").Logger()

	return &logger
}
