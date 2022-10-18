package service

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/djokcik/gophkeeper/client"
	"github.com/djokcik/gophkeeper/client/storage"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/djokcik/gophkeeper/pkg/common"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/rs/zerolog"
	"net/rpc"
)

const (
	CallSaveRecordPersonalDataHandler        = "RPCHandler.SaveRecordPersonalDataHandler"
	CallLoadRecordPersonalDataByKeyHandler   = "RPCHandler.LoadRecordPersonalDataByKeyHandler"
	CallRemoveRecordPersonalDataByKeyHandler = "RPCHandler.RemoveRecordPersonalDataByKeyHandler"

	CallSaveRecordBankCardHandler        = "RPCHandler.SaveRecordBankCardHandler"
	CallLoadRecordBankCardByKeyHandler   = "RPCHandler.LoadRecordBankCardByKeyHandler"
	CallRemoveRecordBankCardByKeyHandler = "RPCHandler.RemoveRecordBankCardByKeyHandler"

	CallSaveRecordTextDataHandler        = "RPCHandler.SaveRecordTextDataHandler"
	CallLoadRecordTextDataByKeyHandler   = "RPCHandler.LoadRecordTextDataByKeyHandler"
	CallRemoveRecordTextDataByKeyHandler = "RPCHandler.RemoveRecordTextDataByKeyHandler"

	CallSaveRecordBinaryDataHandler        = "RPCHandler.SaveRecordBinaryDataHandler"
	CallLoadRecordBinaryDataByKeyHandler   = "RPCHandler.LoadRecordBinaryDataByKeyHandler"
	CallRemoveRecordBinaryDataByKeyHandler = "RPCHandler.RemoveRecordBinaryDataByKeyHandler"

	CallRegisterHandler = "RPCHandler.RegisterHandler"
	CallSignInHandler   = "RPCHandler.SignInHandler"
)

//go:generate mockery --name=ClientRPCService --with-expecter

// ClientRPCService provides methods for call RPC API
type ClientRPCService interface {
	Login(ctx context.Context, username string, password string) (string, error)
	Register(ctx context.Context, username string, password string) (string, error)

	SaveRecordPersonalData(ctx context.Context, token string, key string, data string) error
	LoadRecordPersonalDataByKey(ctx context.Context, token string, key string) (string, error)
	RemoveRecordPersonalDataByKey(ctx context.Context, token string, key string) error

	RemoveRecordBankCardByKey(ctx context.Context, token string, key string) error
	LoadRecordBankCardByKey(ctx context.Context, token string, key string) (string, error)
	SaveRecordBankCard(ctx context.Context, token string, key string, data string) error

	RemoveRecordTextDataByKey(ctx context.Context, token string, key string) error
	LoadRecordTextDataByKey(ctx context.Context, token string, key string) (string, error)
	SaveRecordTextData(ctx context.Context, token string, key string, data string) error

	RemoveRecordBinaryDataByKey(ctx context.Context, token string, key string) error
	LoadRecordBinaryDataByKey(ctx context.Context, token string, key string) (string, error)
	SaveRecordBinaryData(ctx context.Context, token string, key string, data string) error

	CheckOnline() bool
	Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error
}

type rpcService struct {
	cfg       client.Config
	api       *rpc.Client
	crypto    common.CryptoService
	sslConfig common.SSLConfigService

	localStorage storage.ClientLocalStorage
}

func NewRPCService(
	cfg client.Config,
	crypto common.CryptoService,
	sslConfig common.SSLConfigService,
	localStorage storage.ClientLocalStorage,
) ClientRPCService {
	ctx := context.Background()
	service := rpcService{
		cfg:          cfg,
		crypto:       crypto,
		sslConfig:    sslConfig,
		localStorage: localStorage,
	}

	err := service.Reconnect(ctx)
	if err != nil {
		service.Log(ctx).Error().Err(err).Msg("unable to connect to server")
	}

	return &service
}

// Reconnect try to reconnect RPC API
func (s *rpcService) Reconnect(ctx context.Context) error {
	conf, err := s.sslConfig.LoadClientCertificate(s.cfg)
	if err != nil {
		s.api = nil

		s.Log(ctx).Error().Err(err).Msg("Reconnect: err load certificate")
		return ErrUnableConnectServer
	}

	conn, err := tls.Dial("tcp", s.cfg.Address, conf)
	if err != nil {
		s.api = nil

		s.Log(ctx).Error().Err(err).Msgf("Reconnect: invalid to connect tcp server - %s", s.cfg.Address)
		return ErrUnableConnectServer
	}

	s.api = rpc.NewClient(conn)

	return nil
}

// CheckOnline returns true if api exists
func (s rpcService) CheckOnline() bool {
	return s.api != nil
}

// Call should call RPC API with method, args and reply
func (s *rpcService) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	s.Log(ctx).Trace().Msgf("Call: start method - %s, args - %s", serviceMethod, args)

	err := s.api.Call(serviceMethod, args, reply)
	if errors.Is(err, rpc.ErrShutdown) {
		err = s.Reconnect(ctx)
		if err != nil {
			s.Log(ctx).Warn().Err(err).Msg("Call: error reconnect")
			return err
		}

		err = s.api.Call(serviceMethod, args, reply)
	}

	return err
}

// Login call RPC API Login user
func (s rpcService) Login(ctx context.Context, username string, password string) (string, error) {
	if !s.CheckOnline() {
		return "", ErrAnonymousUser
	}

	loginDto := rpcdto.LoginDto{Login: username, Password: password}

	var token string

	err := s.Call(ctx, CallSignInHandler, loginDto, &token)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("SignIn: error call - %s", CallSignInHandler)
		return "", err
	}

	s.Log(ctx).Trace().Msgf("SignIn: user - %s", username)

	return token, nil
}

// Register call RPC API Register user
func (s rpcService) Register(ctx context.Context, username string, password string) (string, error) {
	if !s.CheckOnline() {
		return "", ErrAnonymousUser
	}

	registerDto := rpcdto.RegisterDto{Login: username, Password: password}

	var token string

	err := s.Call(ctx, CallRegisterHandler, registerDto, &token)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("CreateUser: error call - %s", CallRegisterHandler)
		return "", err
	}

	return token, nil
}

// SaveRecordPersonalData call RPC API Save Personal Data Record
func (s rpcService) SaveRecordPersonalData(ctx context.Context, token string, key string, data string) error {
	return s.saveRecord(ctx, rpcdto.SaveRecordRequestDto{Token: token, Key: key, Data: data}, CallSaveRecordPersonalDataHandler)
}

// LoadRecordPersonalDataByKey call RPC API Load Personal Data Record
func (s rpcService) LoadRecordPersonalDataByKey(ctx context.Context, token string, key string) (string, error) {
	return s.loadRecordByKey(ctx, rpcdto.LoadRecordRequestDto{Key: key, Token: token}, CallLoadRecordPersonalDataByKeyHandler)
}

// RemoveRecordPersonalDataByKey call RPC API Remove Personal Data Record
func (s rpcService) RemoveRecordPersonalDataByKey(ctx context.Context, token string, key string) error {
	return s.removeRecordByKey(ctx, rpcdto.RemoveRecordRequestDto{Key: key, Token: token}, CallRemoveRecordPersonalDataByKeyHandler)
}

// SaveRecordBankCard call RPC API Save Personal Bank Card
func (s rpcService) SaveRecordBankCard(ctx context.Context, token string, key string, data string) error {
	return s.saveRecord(ctx, rpcdto.SaveRecordRequestDto{Token: token, Key: key, Data: data}, CallSaveRecordBankCardHandler)
}

// LoadRecordBankCardByKey call RPC API Load Bank Card
func (s rpcService) LoadRecordBankCardByKey(ctx context.Context, token string, key string) (string, error) {
	return s.loadRecordByKey(ctx, rpcdto.LoadRecordRequestDto{Key: key, Token: token}, CallLoadRecordBankCardByKeyHandler)
}

// RemoveRecordBankCardByKey call RPC API Remove Bank Card
func (s rpcService) RemoveRecordBankCardByKey(ctx context.Context, token string, key string) error {
	return s.removeRecordByKey(ctx, rpcdto.RemoveRecordRequestDto{Key: key, Token: token}, CallRemoveRecordBankCardByKeyHandler)
}

// SaveRecordTextData call RPC API Save Text Data
func (s rpcService) SaveRecordTextData(ctx context.Context, token string, key string, data string) error {
	return s.saveRecord(ctx, rpcdto.SaveRecordRequestDto{Token: token, Key: key, Data: data}, CallSaveRecordTextDataHandler)
}

// LoadRecordTextDataByKey call RPC API Load Text Data
func (s rpcService) LoadRecordTextDataByKey(ctx context.Context, token string, key string) (string, error) {
	return s.loadRecordByKey(ctx, rpcdto.LoadRecordRequestDto{Key: key, Token: token}, CallLoadRecordTextDataByKeyHandler)
}

// RemoveRecordTextDataByKey call RPC API Remove Text Data
func (s rpcService) RemoveRecordTextDataByKey(ctx context.Context, token string, key string) error {
	return s.removeRecordByKey(ctx, rpcdto.RemoveRecordRequestDto{Key: key, Token: token}, CallRemoveRecordTextDataByKeyHandler)
}

// SaveRecordBinaryData call RPC API Save Binary Data
func (s rpcService) SaveRecordBinaryData(ctx context.Context, token string, key string, data string) error {
	return s.saveRecord(ctx, rpcdto.SaveRecordRequestDto{Token: token, Key: key, Data: data}, CallSaveRecordBinaryDataHandler)
}

// LoadRecordBinaryDataByKey call RPC API Load Binary Data
func (s rpcService) LoadRecordBinaryDataByKey(ctx context.Context, token string, key string) (string, error) {
	return s.loadRecordByKey(ctx, rpcdto.LoadRecordRequestDto{Key: key, Token: token}, CallLoadRecordBinaryDataByKeyHandler)
}

// RemoveRecordBinaryDataByKey call RPC API Remove Binary Data
func (s rpcService) RemoveRecordBinaryDataByKey(ctx context.Context, token string, key string) error {
	return s.removeRecordByKey(ctx, rpcdto.RemoveRecordRequestDto{Key: key, Token: token}, CallRemoveRecordBinaryDataByKeyHandler)
}

func (s rpcService) loadRecordByKey(ctx context.Context, recordDto rpcdto.LoadRecordRequestDto, serviceMethod string) (string, error) {
	if recordDto.Token == "" {
		return "", ErrAnonymousUser
	}

	var encryptedData string
	err := s.Call(ctx, serviceMethod, recordDto, &encryptedData)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("loadRecordByKey: error call %s", serviceMethod)
		return "", err
	}

	return encryptedData, nil
}

func (s rpcService) removeRecordByKey(ctx context.Context, recordDto rpcdto.RemoveRecordRequestDto, serviceMethod string) error {
	if recordDto.Token == "" {
		return s.ActionRemove(ctx, recordDto.Key, serviceMethod)
	}

	var reply struct{}
	err := s.Call(ctx, serviceMethod, recordDto, &reply)
	if err != nil {
		if errors.Is(err, ErrUnableConnectServer) {
			return s.ActionRemove(ctx, recordDto.Key, serviceMethod)
		}

		s.Log(ctx).Error().Err(err).Msgf("removeRecordByKey: error call %s", serviceMethod)
		return err
	}

	return nil
}

func (s rpcService) saveRecord(ctx context.Context, recordDto rpcdto.SaveRecordRequestDto, serviceMethod string) error {
	if recordDto.Token == "" {
		return s.ActionSave(ctx, recordDto.Key, recordDto.Data, serviceMethod)
	}

	var reply struct{}
	err := s.Call(ctx, serviceMethod, recordDto, &reply)
	if err != nil {
		if errors.Is(err, ErrUnableConnectServer) {
			return s.ActionSave(ctx, recordDto.Key, recordDto.Data, serviceMethod)
		}

		s.Log(ctx).Error().Err(err).Msgf("saveRecord: error call %s", serviceMethod)
		return err
	}

	return nil
}

// ActionRemove mark action as remove and save to local storage
func (s rpcService) ActionRemove(ctx context.Context, key string, method string) error {
	err := s.localStorage.RemoveRecord(ctx, key, method)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("saveRecord: invalid SaveRecord in localStorage")
		return err
	}

	return ErrSaveLocalStorage
}

// ActionSave mark action as save and save to local storage
func (s rpcService) ActionSave(ctx context.Context, key string, data string, method string) error {
	err := s.localStorage.SaveRecord(ctx, key, data, method)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("saveRecord: invalid SaveRecord in localStorage")
		return err
	}

	return ErrSaveLocalStorage
}

func (s rpcService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "RpcService").Logger()

	return &logger
}
