package recordservice

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/client/service"
	"gophkeeper/models"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
)

//go:generate mockery --name=ClientRecordService
type ClientRecordService interface {
	LoadRecordByKey(ctx context.Context, user models.ClientUser, response interface{}, loadFn func() (string, error)) error
	SaveRecord(ctx context.Context, user models.ClientUser, data interface{}, updateFn func(encryptedData string) error) error
}

type recordService struct {
	crypto common.CryptoService
}

func NewClientRecordService(crypto common.CryptoService) ClientRecordService {
	return &recordService{
		crypto: crypto,
	}
}

func (s recordService) LoadRecordByKey(ctx context.Context, user models.ClientUser, response interface{}, loadFn func() (string, error)) error {
	encryptedData, err := loadFn()
	if err != nil {
		s.Log(ctx).Warn().Err(err).Msg("LoadRecordByKey: invalid load data")
		return err
	}

	if encryptedData == "" {
		s.Log(ctx).Trace().Msgf("LoadRecordByKey: data by key not found")
		return service.ErrNotFoundLoadData
	}

	err = s.crypto.DecryptData(ctx, user.Password, encryptedData, &response)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msgf("LoadRecordByKey: invalid decrypt data")
		return err
	}

	s.Log(ctx).Trace().Msg("LoadRecordByKey: load success")

	return err
}

func (s recordService) SaveRecord(ctx context.Context, user models.ClientUser, data interface{}, updateFn func(data string) error) error {
	encryptedData, err := s.crypto.EncryptData(ctx, user.Password, data)
	if err != nil {
		return err
	}

	err = updateFn(encryptedData)
	if err != nil {
		s.Log(ctx).Error().Err(err).Msg("SaveRecord: invalid SaveRecordBankCard")
		return err
	}

	s.Log(ctx).Trace().Msg("LoadRecordByKey: save success")

	return err
}

func (s recordService) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxFileLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "ServerRecordBankCardService").Logger()

	return &logger
}
