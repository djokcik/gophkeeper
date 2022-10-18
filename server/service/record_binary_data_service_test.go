package service

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/server/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_recordBinaryDataService_Save(t *testing.T) {
	t.Run("should save binary data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Save(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, updateFn func(*models.StorageData) error) {
				updateFn(&store)
			}).Return(nil)

		service := recordBinaryDataService{record: &mRecord}
		err := service.Save(ctx, "Data", "username", "binaryData")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{BinaryData: map[string]string{"Data": "binaryData"}})
	})
}

func Test_recordBinaryDataService_Load(t *testing.T) {
	t.Run("should load binary data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{BinaryData: map[string]string{"Data": "binaryData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Load(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, loadFn func(models.StorageData) string) {
				res := loadFn(store)
				require.Equal(t, res, "binaryData")
			}).Return("testReturn", nil)

		service := recordBinaryDataService{record: &mRecord}
		data, err := service.Load(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, data, "testReturn")
	})
}

func Test_recordBinaryDataService_Remove(t *testing.T) {
	t.Run("should remove binary data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{BinaryData: map[string]string{"Data": "binaryData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Remove(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, removeFn func(*models.StorageData) error) {
				removeFn(&store)
			}).Return(nil)

		service := recordBinaryDataService{record: &mRecord}
		err := service.Remove(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{BinaryData: map[string]string{}})
	})
}
