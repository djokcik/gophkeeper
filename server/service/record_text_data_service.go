package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
)

type ServerRecordTextDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
	Remove(ctx context.Context, key string, username string) error
}

type recordTextDataService struct {
	cfg    server.Config
	record ServerRecordService
}

func NewServerRecordTextDataService(cfg server.Config, record ServerRecordService) ServerRecordTextDataService {
	return &recordTextDataService{
		cfg:    cfg,
		record: record,
	}
}

func (s recordTextDataService) Save(ctx context.Context, key string, username string, data string) error {
	return s.record.Save(ctx, username, func(store models.StorageData) error {
		if store.TextData == nil {
			store.TextData = make(map[string]string)
		}

		store.TextData[key] = data

		return nil
	})
}

func (s recordTextDataService) Load(ctx context.Context, key string, username string) (string, error) {
	return s.record.Load(ctx, username, func(store models.StorageData) string {
		return store.TextData[key]
	})
}

func (s recordTextDataService) Remove(ctx context.Context, key string, username string) error {
	return s.record.Remove(ctx, username, func(store models.StorageData) error {
		if store.TextData == nil {
			return ErrNotFoundRecord
		}

		if _, ok := store.TextData[key]; !ok {
			return ErrNotFoundRecord
		}

		delete(store.TextData, key)

		return nil
	})
}

func (s *recordTextDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordTextDataService").Logger()

	return &logger
}
