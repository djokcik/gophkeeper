package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
)

//go:generate mockery --name=ServerRecordBankCardDataService
type ServerRecordBankCardDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
	Remove(ctx context.Context, key string, username string) error
}

type recordBankCardDataService struct {
	cfg    server.Config
	record ServerRecordService
}

func NewServerRecordBankCardDataService(cfg server.Config, record ServerRecordService) ServerRecordBankCardDataService {
	return &recordBankCardDataService{
		cfg:    cfg,
		record: record,
	}
}

func (s recordBankCardDataService) Save(ctx context.Context, key string, username string, data string) error {
	return s.record.Save(ctx, username, func(store models.StorageData) error {
		if store.BankCardData == nil {
			store.BankCardData = make(map[string]string)
		}

		store.BankCardData[key] = data

		return nil
	})
}

func (s recordBankCardDataService) Load(ctx context.Context, key string, username string) (string, error) {
	return s.record.Load(ctx, username, func(store models.StorageData) string {
		return store.BankCardData[key]
	})
}

func (s recordBankCardDataService) Remove(ctx context.Context, key string, username string) error {
	return s.record.Remove(ctx, username, func(store models.StorageData) error {
		if store.BankCardData == nil {
			return ErrNotFoundRecord
		}

		if _, ok := store.BankCardData[key]; !ok {
			return ErrNotFoundRecord
		}

		delete(store.BankCardData, key)

		return nil
	})
}

func (s *recordBankCardDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBankCardDataService").Logger()

	return &logger
}
