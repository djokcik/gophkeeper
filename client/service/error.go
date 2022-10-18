package service

import "errors"

var (
	ErrAnonymousUser       = errors.New("ClientStorageService: user is not set")
	ErrUnableConnectServer = errors.New("ClientRPCService: unable connect to server")
	ErrSaveLocalStorage    = errors.New("ClientRPCService: save to local data because server isn`t connect")

	ErrNotFoundLoadData = errors.New("service: err not found load data")
	ErrInvalidLoadData  = errors.New("service: err invalid load data")
)
