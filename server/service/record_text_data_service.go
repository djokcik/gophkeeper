package service

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=ServerRecordTextDataService --with-expecter

// ServerRecordTextDataService provide methods for control TextData
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

// Save is save record text data
func (s recordTextDataService) Save(ctx context.Context, key string, username string, data string) error {
	return s.record.Save(ctx, username, func(store *models.StorageData) error {
		if store.TextData == nil {
			store.TextData = make(map[string]string)
		}

		store.TextData[key] = data

		return nil
	})
}

// Load is load record text data
func (s recordTextDataService) Load(ctx context.Context, key string, username string) (string, error) {
	return s.record.Load(ctx, username, func(store models.StorageData) string {
		return store.TextData[key]
	})
}

// Remove is remove record text data
func (s recordTextDataService) Remove(ctx context.Context, key string, username string) error {
	return s.record.Remove(ctx, username, func(store *models.StorageData) error {
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
