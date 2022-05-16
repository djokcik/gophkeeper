package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/models"
	"gophkeeper/server/service/mocks"
	"testing"
)

func Test_recordPersonalDataService_Save(t *testing.T) {
	t.Run("should save personal data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Save(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, updateFn func(*models.StorageData) error) {
				updateFn(&store)
			}).Return(nil)

		service := recordPersonalDataService{record: &mRecord}
		err := service.Save(ctx, "Data", "username", "personalData")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{PersonalData: map[string]string{"Data": "personalData"}})
	})
}

func Test_recordPersonalDataService_Load(t *testing.T) {
	t.Run("should load personal data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{PersonalData: map[string]string{"Data": "personalData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Load(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, loadFn func(models.StorageData) string) {
				res := loadFn(store)
				require.Equal(t, res, "personalData")
			}).Return("testReturn", nil)

		service := recordPersonalDataService{record: &mRecord}
		data, err := service.Load(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, data, "testReturn")
	})
}

func Test_recordPersonalDataService_Remove(t *testing.T) {
	t.Run("should remove personal data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{PersonalData: map[string]string{"Data": "personalData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Remove(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, removeFn func(*models.StorageData) error) {
				removeFn(&store)
			}).Return(nil)

		service := recordPersonalDataService{record: &mRecord}
		err := service.Remove(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{PersonalData: map[string]string{}})
	})
}
