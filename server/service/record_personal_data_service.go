package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
)

type ServerRecordPersonalDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
	Remove(ctx context.Context, key string, username string) error
}

type recordPersonalDataService struct {
	cfg    server.Config
	record ServerRecordService
}

func NewServerRecordPersonalDataService(cfg server.Config, record ServerRecordService) ServerRecordPersonalDataService {
	return &recordPersonalDataService{
		cfg:    cfg,
		record: record,
	}
}

func (s recordPersonalDataService) Save(ctx context.Context, key string, username string, data string) error {
	return s.record.Save(ctx, username, func(store models.StorageData) error {
		if store.PersonalData == nil {
			store.PersonalData = make(map[string]string)
		}

		store.PersonalData[key] = data

		return nil
	})
}

func (s recordPersonalDataService) Load(ctx context.Context, key string, username string) (string, error) {
	return s.record.Load(ctx, username, func(store models.StorageData) string {
		return store.PersonalData[key]
	})
}

func (s recordPersonalDataService) Remove(ctx context.Context, key string, username string) error {
	return s.record.Remove(ctx, username, func(store models.StorageData) error {
		if store.PersonalData == nil {
			return ErrNotFoundRecord
		}

		if _, ok := store.PersonalData[key]; !ok {
			return ErrNotFoundRecord
		}

		delete(store.PersonalData, key)

		return nil
	})
}

func (s *recordPersonalDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordPersonalDataService").Logger()

	return &logger
}
