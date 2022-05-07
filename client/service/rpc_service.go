package service

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/rs/zerolog"
	"gophkeeper/client"
	"gophkeeper/client/storage"
	"gophkeeper/models/rpcdto"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
	"net/rpc"
)

const (
	CallSaveRecordPersonalDataHandler     = "RpcHandler.SaveRecordPersonalDataHandler"
	CallLoadRecordPrivateDataByKeyHandler = "RpcHandler.LoadRecordPrivateDataByKeyHandler"
	CallRegisterHandler                   = "RpcHandler.RegisterHandler"
	CallSignInHandler                     = "RpcHandler.SignInHandler"
)

type ClientRpcService interface {
	Login(ctx context.Context, username string, password string) (string, error)
	Register(ctx context.Context, username string, password string) (string, error)

	SaveRecordPersonalData(ctx context.Context, token string, key string, data string) error
	LoadRecordPrivateDataByKey(ctx context.Context, token string, key string) (string, error)

	CheckOnline() bool
	Call(ctx context.Context, serviceMethod string, args any, reply any) error
}

type rpcService struct {
	cfg       client.Config
	api       *rpc.Client
	crypto    common.CryptoService
	sslConfig common.SSLConfigService

	localStorage storage.ClientLocalStorage
}

func NewRpcService(
	cfg client.Config,
	crypto common.CryptoService,
	sslConfig common.SSLConfigService,
	localStorage storage.ClientLocalStorage,
) ClientRpcService {
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

func (s rpcService) CheckOnline() bool {
	return s.api != nil
}

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

func (s rpcService) SaveRecordPersonalData(ctx context.Context, token string, key string, data string) error {
	return s.saveRecord(ctx, rpcdto.SaveRecordRequestDto{Token: token, Key: key, Data: data}, CallSaveRecordPersonalDataHandler)
}

func (s rpcService) LoadRecordPrivateDataByKey(ctx context.Context, token string, key string) (string, error) {
	return s.loadRecordByKey(ctx, rpcdto.LoadRecordRequestDto{Key: key, Token: token}, CallLoadRecordPrivateDataByKeyHandler)
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

func (s rpcService) saveRecord(ctx context.Context, recordDto rpcdto.SaveRecordRequestDto, serviceMethod string) error {
	if recordDto.Token == "" {
		return s.SaveAction(ctx, recordDto.Key, recordDto.Data, serviceMethod)
	}

	var reply struct{}
	err := s.Call(ctx, serviceMethod, recordDto, &reply)
	if err != nil {
		if errors.Is(err, ErrUnableConnectServer) {
			return s.SaveAction(ctx, recordDto.Key, recordDto.Data, serviceMethod)
		}

		s.Log(ctx).Error().Err(err).Msgf("saveRecord: error call %s", serviceMethod)
		return err
	}

	return nil
}

func (s rpcService) SaveAction(ctx context.Context, key string, data string, method string) error {
	err := s.localStorage.SaveRecord(ctx, key, data, method)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("saveRecord: invalid SaveRecord in localStorage")
		return err
	}

	return ErrSaveLocalStorage
}

func (s rpcService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "rpcService").Logger()

	return &logger
}
