package storage

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"os"
)

type FileStorage interface {
	Read(ctx context.Context, filepath string, in interface{}) error
	Write(ctx context.Context, filepath string, out interface{}) error
}

type fileStorage struct {
	Cfg server.Config
}

func NewFileStorage(cfg server.Config) FileStorage {
	return &fileStorage{Cfg: cfg}
}

func (s fileStorage) Read(ctx context.Context, filepath string, in interface{}) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Read: err open file")
		return err
	}

	err = json.NewDecoder(file).Decode(in)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("err decode json in file")
		return err
	}

	return nil
}

func (s fileStorage) Write(ctx context.Context, filepath string, out interface{}) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Write: err open file")
		return err
	}

	err = json.NewEncoder(file).Encode(out)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("Write: err write to file")
		return err
	}

	return nil
}

func (s *fileStorage) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "FileStorage").Logger()

	return &logger
}
