package service

import "errors"

var (
	ErrUnableConnectServer = errors.New("ClientRpcService: unable connect to server")
	ErrNotFoundLoadData    = errors.New("view: err not found load data")
)
