package storage

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/clientmodels"
	"gophkeeper/pkg/logging"
	"os"
	"path"
)

type ClientLocalStorage interface {
	ClearActions(ctx context.Context) error
	SaveRecord(ctx context.Context, key string, data string, method string) error
	LoadRecords(ctx context.Context) ([]clientmodels.RecordFileLine, error)
	Close()
}

type localStorage struct {
	FileReader *actionFileStoreReader
	FileWriter *actionFileStoreWriter
}

func NewClientLocalStorage(ctx context.Context) ClientLocalStorage {
	s := &localStorage{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		logging.NewFileLogger().Fatal().Err(err).Msg("NewClientLocalStorage: err get UserHomeDir")
		return nil
	}

	dir := path.Join(homeDir, ".gk")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		s.Log(ctx).Fatal().Err(err).Msg("NewClientLocalStorage: can`t make directory")
		return nil
	}

	filepath := path.Join(dir, "actions.txt")

	reader, err := newActionFileStoreReader(filepath)
	if err != nil {
		s.Log(ctx).Fatal().Err(err).Msg("NewClientLocalStorage: failed to open reader")
		return nil
	}

	writer, err := newActionFileStoreWriter(filepath)
	if err != nil {
		s.Log(ctx).Fatal().Err(err).Msg("NewClientLocalStorage: failed to open writer")
		return nil
	}

	s.FileReader = reader
	s.FileWriter = writer

	return s
}

func (s localStorage) ClearActions(_ context.Context) error {
	return s.FileWriter.SaveActions(storeActions{})
}

func (s localStorage) LoadRecords(_ context.Context) ([]clientmodels.RecordFileLine, error) {
	return s.FileReader.ReadActions()
}

func (s localStorage) SaveRecord(ctx context.Context, key string, data string, method string) error {
	fileLine := clientmodels.RecordFileLine{
		Data:       data,
		Key:        key,
		Method:     method,
		ActionType: clientmodels.SaveMethod,
	}

	actions, err := s.FileReader.ReadActions()
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: err read actions")
		return err
	}

	actions = append(actions, fileLine)

	err = s.FileWriter.SaveActions(actions)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: err save actions")
		return err
	}

	return nil
}

func (s *localStorage) Close() {
	s.FileReader.Close()
	s.FileWriter.Close()
}

func (s localStorage) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "LocalStorage").Logger()

	return &logger
}
