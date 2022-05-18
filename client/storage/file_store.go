package storage

import (
	"encoding/json"
	"errors"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	"io"
	"os"
)

//go:generate mockery --name=ActionFileStoreReader --with-expecter

// ActionFileStoreReader provides methods for read actions
type ActionFileStoreReader interface {
	ReadActions() (clientmodels.StoreActions, error)
	Close() error
}

type actionFileStoreReader struct {
	file    *os.File
	decoder *json.Decoder
}

func newActionFileStoreReader(filename string) (ActionFileStoreReader, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	return &actionFileStoreReader{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

// ReadActions read list of actions
func (r *actionFileStoreReader) ReadActions() (clientmodels.StoreActions, error) {
	_, err := r.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	var actions clientmodels.StoreActions
	err = r.decoder.Decode(&actions)

	if errors.Is(err, io.EOF) {
		return []clientmodels.RecordFileLine{}, nil
	}

	return actions, err
}

func (r *actionFileStoreReader) Close() error {
	return r.file.Close()
}

//go:generate mockery --name=ActionFileStoreWriter --with-expecter

// ActionFileStoreWriter provides methods for write actions
type ActionFileStoreWriter interface {
	SaveActions(actions clientmodels.StoreActions) error
	Close() error
}

type actionFileStoreWriter struct {
	file    *os.File
	encoder *json.Encoder
}

func newActionFileStoreWriter(filename string) (ActionFileStoreWriter, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	return &actionFileStoreWriter{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

// SaveActions save to file list of actions
func (w *actionFileStoreWriter) SaveActions(actions clientmodels.StoreActions) error {
	err := w.file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = w.file.Seek(0, 0)
	if err != nil {
		return err
	}

	return w.encoder.Encode(actions)
}

func (w *actionFileStoreWriter) Close() error {
	return w.file.Close()
}
