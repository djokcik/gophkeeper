package service

import "errors"

var (
	ErrAnonymousUser       = errors.New("ClientStorageService: user is not set")
	ErrUnableConnectServer = errors.New("ClientRpcService: unable connect to server")
	ErrSaveLocalStorage    = errors.New("ClientRpcService: save to local data because server isn`t connect")

	ErrNotFoundLoadData = errors.New("service: err not found load data")
)
