package storage

import (
	"encoding/json"
	"errors"
	"gophkeeper/client/clientmodels"
	"io"
	"os"
)

type storeActions []clientmodels.RecordFileLine

type actionFileStoreReader struct {
	file    *os.File
	decoder *json.Decoder
}

func newActionFileStoreReader(filename string) (*actionFileStoreReader, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	return &actionFileStoreReader{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (r *actionFileStoreReader) ReadActions() (storeActions, error) {
	_, err := r.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	var actions storeActions
	err = r.decoder.Decode(&actions)

	if errors.Is(err, io.EOF) {
		return []clientmodels.RecordFileLine{}, nil
	}

	return actions, err
}

func (r *actionFileStoreReader) Close() error {
	return r.file.Close()
}

type actionFileStoreWriter struct {
	file    *os.File
	encoder *json.Encoder
}

func newActionFileStoreWriter(filename string) (*actionFileStoreWriter, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	return &actionFileStoreWriter{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (w *actionFileStoreWriter) SaveActions(actions storeActions) error {
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
