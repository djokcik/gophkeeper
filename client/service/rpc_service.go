package service

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"gophkeeper/client"
	"gophkeeper/models"
	"gophkeeper/models/rpcdto"
	"gophkeeper/pkg/logging"
	"net/rpc"
)

type RpcService interface {
	Login(ctx context.Context, username string, password string) (models.GophUser, error)
	Register(ctx context.Context, username string, password string) (models.GophUser, error)

	SaveLoginPassword(ctx context.Context, user models.GophUser, username string, password string) error
	LoadPasswordByLogin(ctx context.Context, user models.GophUser, username string) (models.LoginPasswordResponseDto, error)
}

type rpcService struct {
	cfg  client.Config
	call *rpc.Client
}

func NewRpcService(cfg client.Config) RpcService {
	service := rpcService{cfg: cfg}

	service.Reconnect()

	return &service
}

func (s *rpcService) Reconnect() {
	rpcClient, err := rpc.Dial("tcp", s.cfg.Address)
	if err != nil {
		logging.NewFileLogger().Fatal().Err(err).Msgf("invalid to connect tcp server - %s", s.cfg.Address)
	}

	s.call = rpcClient
}

func (s *rpcService) Call(serviceMethod string, args any, reply any) error {
	err := s.call.Call(serviceMethod, args, reply)
	if errors.Is(err, rpc.ErrShutdown) {
		s.Reconnect()
		err = s.call.Call(serviceMethod, args, reply)
	}

	return err
}

func (s rpcService) Login(ctx context.Context, username string, password string) (models.GophUser, error) {
	loginDto := rpcdto.LoginDto{Login: username, Password: password}

	var user models.GophUser

	err := s.Call("RpcHandler.LoginHandler", loginDto, &user)

	s.Log(ctx).Trace().Msgf("Login: user - %v", user)

	return user, err
}

func (s rpcService) Register(_ context.Context, username string, password string) (models.GophUser, error) {
	registerDto := rpcdto.RegisterDto{Login: username, Password: password}

	var user models.GophUser

	err := s.Call("RpcHandler.RegisterHandler", registerDto, &user)

	return user, err
}

func (s rpcService) SaveLoginPassword(ctx context.Context, user models.GophUser, username string, password string) error {
	loginDto := rpcdto.SaveLoginPasswordDto{Login: username, Password: password, User: user}

	err := s.Call("RpcHandler.SaveLoginPasswordHandler", loginDto, &struct{}{})

	s.Log(ctx).Trace().Msg("SaveLoginPassword: data")

	return err
}

func (s rpcService) LoadPasswordByLogin(ctx context.Context, user models.GophUser, username string) (models.LoginPasswordResponseDto, error) {
	loadDto := rpcdto.LoadLoginPasswordDto{Login: username, User: user}

	var data models.LoginPasswordResponseDto

	err := s.Call("RpcHandler.LoadPasswordByLoginHandler", loadDto, &data)

	s.Log(ctx).Trace().Msg("SaveLoginPassword: data")

	return data, err
}

func (s rpcService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "rpcService").Logger()

	return &logger
}
