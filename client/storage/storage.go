package storage

import (
	"context"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/rs/zerolog"
	"os"
	"path"
)

//go:generate mockery --name=ClientLocalStorage --with-expecter

// ClientLocalStorage provides methods for control client LocalStorage
type ClientLocalStorage interface {
	ClearActions(ctx context.Context) error
	SaveRecord(ctx context.Context, key string, data string, method string) error
	RemoveRecord(ctx context.Context, key string, method string) error
	LoadRecords(ctx context.Context) ([]clientmodels.RecordFileLine, error)
	Close()
}

type localStorage struct {
	FileReader ActionFileStoreReader
	FileWriter ActionFileStoreWriter
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

// ClearActions clear list of actions
func (s localStorage) ClearActions(_ context.Context) error {
	return s.FileWriter.SaveActions(clientmodels.StoreActions{})
}

// LoadRecords returns list of records
func (s localStorage) LoadRecords(_ context.Context) ([]clientmodels.RecordFileLine, error) {
	return s.FileReader.ReadActions()
}

// RemoveRecord add new record with type remove
func (s localStorage) RemoveRecord(ctx context.Context, key string, method string) error {
	fileLine := clientmodels.RecordFileLine{
		Key:        key,
		Method:     method,
		ActionType: clientmodels.RemoveMethod,
	}

	actions, err := s.FileReader.ReadActions()
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("RemoveRecord: err read actions")
		return err
	}

	actions = append(actions, fileLine)

	err = s.FileWriter.SaveActions(actions)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("RemoveRecord: err save actions")
		return err
	}

	return nil
}

// SaveRecord add new record with type save
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
