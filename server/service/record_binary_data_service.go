package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/models"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
)

//go:generate mockery --name=ServerRecordBinaryDataService  --with-expecter
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

func (s recordBinaryDataService) Save(ctx context.Context, key string, username string, data string) error {
	return s.record.Save(ctx, username, func(store *models.StorageData) error {
		if store.BinaryData == nil {
			store.BinaryData = make(map[string]string)
		}

		store.BinaryData[key] = data

		return nil
	})
}

func (s recordBinaryDataService) Load(ctx context.Context, key string, username string) (string, error) {
	return s.record.Load(ctx, username, func(store models.StorageData) string {
		return store.BinaryData[key]
	})
}

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
