package service

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/server/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_recordTextDataService_Save(t *testing.T) {
	t.Run("should save text data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.On("Save", ctx, "username", mock.Anything).Run(func(args mock.Arguments) {
			fn := args.Get(2).(func(store *models.StorageData) error)
			fn(&store)
		}).Return(nil)

		service := recordTextDataService{record: &mRecord}
		err := service.Save(ctx, "Data", "username", "textData")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{TextData: map[string]string{"Data": "textData"}})
	})
}

func Test_recordTextDataService_Load(t *testing.T) {
	t.Run("should load text data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{TextData: map[string]string{"Data": "textData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().Load(ctx, "username", mock.Anything).
			Run(func(ctx context.Context, username string, loadFn func(models.StorageData) string) {
				res := loadFn(store)
				require.Equal(t, res, "textData")
			}).Return("testReturn", nil)

		service := recordTextDataService{record: &mRecord}
		data, err := service.Load(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, data, "testReturn")
	})
}

func Test_recordTextDataService_Remove(t *testing.T) {
	t.Run("should remove text data", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{TextData: map[string]string{"Data": "textData"}}

		mRecord := mocks.ServerRecordService{Mock: mock.Mock{}}
		mRecord.On("Remove", ctx, "username", mock.Anything).Run(func(args mock.Arguments) {
			fn := args.Get(2).(func(store *models.StorageData) error)
			fn(&store)
		}).Return(nil)

		service := recordTextDataService{record: &mRecord}
		err := service.Remove(ctx, "Data", "username")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{TextData: map[string]string{}})
	})
}
