package service

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server"
	"github.com/rs/zerolog"
)

//go:generate mockery --name=ServerRecordBinaryDataService  --with-expecter

// ServerRecordBinaryDataService provide methods for control BinaryData
type ServerRecordBinaryDataService interface {
	Save(ctx context.Context, key string, username string, data string) error
	Load(ctx context.Context, key string, username string) (string, error)
	Remove(ctx context.Context, key string, username string) error
}

type recordBinaryDataService struct {
	cfg    server.Config
	record ServerRecordService
}

func NewServerRecordBinaryDataService(cfg server.Config, record ServerRecordService) ServerRecordBinaryDataService {
	return &recordBinaryDataService{
		cfg:    cfg,
		record: record,
	}
}

// Save is save record binary data
func (s recordBinaryDataService) Save(ctx context.Context, key string, username string, data string) error {
	return s.record.Save(ctx, username, func(store *models.StorageData) error {
		if store.BinaryData == nil {
			store.BinaryData = make(map[string]string)
		}

		store.BinaryData[key] = data

		return nil
	})
}

// Load is load record binary data
func (s recordBinaryDataService) Load(ctx context.Context, key string, username string) (string, error) {
	return s.record.Load(ctx, username, func(store models.StorageData) string {
		return store.BinaryData[key]
	})
}

// Remove is remove record binary data
func (s recordBinaryDataService) Remove(ctx context.Context, key string, username string) error {
	return s.record.Remove(ctx, username, func(store *models.StorageData) error {
		if store.BinaryData == nil {
			return ErrNotFoundRecord
		}

		if _, ok := store.BinaryData[key]; !ok {
			return ErrNotFoundRecord
		}

		delete(store.BinaryData, key)

		return nil
	})
}

func (s *recordBinaryDataService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBinaryDataService").Logger()

	return &logger
}
