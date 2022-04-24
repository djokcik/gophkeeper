package service

import (
	"context"
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
}

type rpcService struct {
	call *rpc.Client
}

func NewRpcService(cfg client.Config) RpcService {
	rpcClient, err := rpc.Dial("tcp", cfg.Address)
	if err != nil {
		logging.NewFileLogger().Fatal().Err(err).Msgf("invalid to connect tcp server - %s", cfg.Address)
	}

	return &rpcService{
		call: rpcClient,
	}
}

func (s rpcService) Login(ctx context.Context, username string, password string) (models.GophUser, error) {
	loginDto := rpcdto.LoginDto{Login: username, Password: password}

	var user models.GophUser

	err := s.call.Call("RpcHandler.Login", loginDto, &user)

	s.Log(ctx).Trace().Msgf("Login: user - %v", user)

	return user, err
}

func (s rpcService) Register(_ context.Context, username string, password string) (models.GophUser, error) {
	registerDto := rpcdto.RegisterDto{Login: username, Password: password}

	var user models.GophUser

	err := s.call.Call("RpcHandler.Register", registerDto, &user)

	return user, err
}

func (s rpcService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "rpcService").Logger()

	return &logger
}
