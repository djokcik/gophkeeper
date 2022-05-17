package service

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/server/service/mocks"
	mocks2 "github.com/djokcik/gophkeeper/server/storage/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_recordService_Save(t *testing.T) {
	t.Run("should save record", func(t *testing.T) {
		ctx := context.TODO()

		mLock := mocks.KeyLockService{Mock: mock.Mock{}}
		mLock.On("Lock", "username")
		mLock.On("Unlock", "username")

		store := models.StorageData{TextData: map[string]string{"Data": "textData"}}

		mStorage := mocks2.Storage{Mock: mock.Mock{}}
		mStorage.On("Read", ctx, "username").Return(store, nil)
		mStorage.On("Save", ctx, store).Return(nil)

		service := recordService{keyLock: &mLock, storage: &mStorage}

		flag := false
		err := service.Save(ctx, "username", func(store *models.StorageData) error {
			flag = true
			return nil
		})

		require.Equal(t, flag, true)
		require.Equal(t, err, nil)
	})
}

func Test_recordService_Load(t *testing.T) {
	t.Run("should load record", func(t *testing.T) {
		ctx := context.TODO()

		store := models.StorageData{TextData: map[string]string{"Data": "textData"}}

		mStorage := mocks2.Storage{Mock: mock.Mock{}}
		mStorage.On("Read", ctx, "username").Return(store, nil)

		service := recordService{storage: &mStorage}

		resp, err := service.Load(ctx, "username", func(store models.StorageData) string {
			return store.TextData["Data"]
		})

		require.Equal(t, resp, "textData")
		require.Equal(t, err, nil)
	})
}

func Test_recordService_Remove(t *testing.T) {
	t.Run("should remove record", func(t *testing.T) {
		ctx := context.TODO()

		mLock := mocks.KeyLockService{Mock: mock.Mock{}}
		mLock.On("Lock", "username")
		mLock.On("Unlock", "username")

		store := models.StorageData{TextData: map[string]string{"Data": "textData"}}

		mStorage := mocks2.Storage{Mock: mock.Mock{}}
		mStorage.On("Read", ctx, "username").Return(store, nil)
		mStorage.On("Save", ctx, store).Return(nil)

		service := recordService{keyLock: &mLock, storage: &mStorage}

		flag := false
		err := service.Remove(ctx, "username", func(store *models.StorageData) error {
			flag = true
			return nil
		})

		require.Equal(t, flag, true)
		require.Equal(t, err, nil)
	})
}
