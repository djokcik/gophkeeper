package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/models"
	"gophkeeper/server/service/mocks"
	"testing"
)

func Test_recordBankCardDataService_Save(t *testing.T) {
	t.Run("should save bank card data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Save(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, updateFn func(*models.StorageData) error) {
				updateFn(&store)
			}).Return(nil)

		service := recordBankCardDataService{record: &mRecord}
		err := service.Save(ctx, "Data", "username", "bankCardData")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{BankCardData: map[string]string{"Data": "bankCardData"}})
	})
}

func Test_recordBankCardDataService_Load(t *testing.T) {
	t.Run("should load bank card data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{BankCardData: map[string]string{"Data": "bankCardData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Load(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, loadFn func(models.StorageData) string) {
				res := loadFn(store)
				require.Equal(t, res, "bankCardData")
			}).Return("testReturn", nil)

		service := recordBankCardDataService{record: &mRecord}
		data, err := service.Load(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, data, "testReturn")
	})
}

func Test_recordBankCardDataService_Remove(t *testing.T) {
	t.Run("should remove bank card data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{BankCardData: map[string]string{"Data": "bankCardData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Remove(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, removeFn func(*models.StorageData) error) {
				removeFn(&store)
			}).Return(nil)

		service := recordBankCardDataService{record: &mRecord}
		err := service.Remove(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{BankCardData: map[string]string{}})
	})
}
