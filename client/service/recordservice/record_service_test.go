package recordservice

import (
	"context"
	"encoding/json"
	service2 "github.com/djokcik/gophkeeper/client/service"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/common/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type testStruct struct {
	Data string `json:"data"`
}

func Test_recordService_LoadRecordByKey(t *testing.T) {
	t.Run("should return error when data isn`t found", func(t *testing.T) {
		ctx := context.TODO()

		service := recordService{}

		flag := false
		err := service.LoadRecordByKey(ctx, models.ClientUser{}, nil, func() (string, error) {
			flag = true
			return "", nil
		})

		require.Equal(t, err, service2.ErrNotFoundLoadData)
		require.Equal(t, flag, true)
	})

	t.Run("should load decrypted data", func(t *testing.T) {
		ctx := context.TODO()

		var response testStruct

		mCrypto := mocks.CryptoService{Mock: mock.Mock{}}
		mCrypto.EXPECT().DecryptData(ctx, "password", "encryptData", mock.Anything).
			Run(func(ctx context.Context, userPassword string, encryptedData string, res interface{}) {
				err := json.Unmarshal([]byte(`{"data":"testData"}`), &res)
				require.Equal(t, err, nil)
			}).Return(nil)

		service := recordService{crypto: &mCrypto}

		flag := false
		err := service.LoadRecordByKey(ctx, models.ClientUser{Password: "password"}, &response, func() (string, error) {
			flag = true
			return "encryptData", nil
		})

		require.Equal(t, err, nil)
		require.Equal(t, flag, true)
		require.Equal(t, response, testStruct{Data: "testData"})
	})
}

func Test_recordService_SaveRecord(t *testing.T) {
	t.Run("record should be saved", func(t *testing.T) {
		ctx := context.TODO()

		data := testStruct{Data: "test"}

		mCrypto := mocks.CryptoService{Mock: mock.Mock{}}
		mCrypto.EXPECT().EncryptData(ctx, "password", data).Return("testData", nil)

		svc := recordService{crypto: &mCrypto}
		flag := false
		err := svc.SaveRecord(ctx, models.ClientUser{Password: "password"}, data, func(data string) error {
			require.Equal(t, data, "testData")
			flag = true
			return nil
		})

		require.Equal(t, err, nil)
		require.Equal(t, flag, true)
	})
}
