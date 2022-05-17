package service

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/storage"
	"gophkeeper/models/rpcdto"
	"gophkeeper/pkg/logging"
)

//go:generate mockery --name=ClientStorageService --with-expecter
type ClientStorageService interface {
	LoadRecords(ctx context.Context) ([]clientmodels.RecordFileLine, error)
	SyncServer(ctx context.Context) error
}

type storageService struct {
	api  ClientRpcService
	user ClientUserService

	store storage.ClientLocalStorage
}

func NewClientStorageService(
	api ClientRpcService,
	localStorage storage.ClientLocalStorage,
	user ClientUserService,
) ClientStorageService {
	return &storageService{
		api:   api,
		store: localStorage,
		user:  user,
	}
}

func (s storageService) SyncServer(ctx context.Context) error {
	actions, err := s.LoadRecords(ctx)
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("SyncServer: invalid load records")
		return err
	}

	s.Log(ctx).Trace().Msgf("SyncServer: get actions len - %d", len(actions))

	user := s.user.GetUser()
	if user.Token == "" {
		s.Log(ctx).Warn().Msgf("SyncServer: user isn`t token")
		return ErrUnableConnectServer
	}

	for _, action := range actions {
		s.Log(ctx).Trace().Msgf("SyncServer: start action - %+v", action)
		switch action.ActionType {
		case clientmodels.SaveMethod:
			var reply struct{}
			err = s.api.Call(ctx, action.Method, rpcdto.SaveRecordRequestDto{Key: action.Key, Data: action.Data, Token: user.Token}, &reply)
			if err != nil {
				s.Log(ctx).Warn().Err(err).Msgf("SyncServer: invalid call save action - %+v", action)
				return err
			}
		case clientmodels.RemoveMethod:
			var reply struct{}
			err = s.api.Call(ctx, action.Method, rpcdto.RemoveRecordRequestDto{Key: action.Key, Token: user.Token}, &reply)
			if err != nil {
				s.Log(ctx).Warn().Err(err).Msgf("SyncServer: invalid call remove action - %+v", action)
				return err
			}
		default:
			s.Log(ctx).Warn().Msgf("unknown action - %+v", action)
		}

		s.Log(ctx).Trace().Msg("SyncServer: end action")
	}

	err = s.store.ClearActions(ctx)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SyncServer: err clear actions")
		return err
	}

	return nil
}

func (s storageService) LoadRecords(ctx context.Context) ([]clientmodels.RecordFileLine, error) {
	return s.store.LoadRecords(ctx)
}

func (s storageService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "StorageService").Logger()

	return &logger
}
