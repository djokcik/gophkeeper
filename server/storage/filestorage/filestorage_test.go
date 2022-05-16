package filestorage

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/models"
	"gophkeeper/server/storage"
	"gophkeeper/server/storage/filestorage/mocks"
	"testing"
)

func Test_fileStorage_Read(t *testing.T) {
	t.Run("should read storage", func(t *testing.T) {
		ctx := context.TODO()

		mBind := mocks.FileBinder{Mock: mock.Mock{}}
		mBind.EXPECT().
			ReadStorage(ctx, "username").Return(models.StorageData{TextData: map[string]string{}}, nil)

		service := fileStorage{bind: &mBind}
		store, err := service.Read(ctx, "username")

		require.Equal(t, err, nil)
		require.Equal(t, store, models.StorageData{TextData: map[string]string{}})
	})
}

func Test_fileStorage_Save(t *testing.T) {
	t.Run("should save to storage", func(t *testing.T) {
		ctx := context.TODO()

		mBind := mocks.FileBinder{Mock: mock.Mock{}}
		mBind.EXPECT().
			SaveStorage(ctx, models.StorageData{TextData: map[string]string{}}).Return(nil)

		service := fileStorage{bind: &mBind}
		err := service.Save(ctx, models.StorageData{TextData: map[string]string{}})

		require.Equal(t, err, nil)
	})
}

func Test_fileStorage_UserByUsername(t *testing.T) {
	t.Run("should return user by username", func(t *testing.T) {
		ctx := context.TODO()

		mBind := mocks.FileBinder{Mock: mock.Mock{}}
		mBind.EXPECT().CheckFileExist(ctx, "username").Return(true, nil)
		mBind.EXPECT().
			ReadStorage(ctx, "username").Return(models.StorageData{User: models.User{Username: "test"}}, nil)

		service := fileStorage{bind: &mBind}
		user, err := service.UserByUsername(ctx, "username")

		require.Equal(t, err, nil)
		require.Equal(t, user, models.User{Username: "test"})
	})
}

func Test_fileStorage_CreateUser(t *testing.T) {
	t.Run("should be created user", func(t *testing.T) {
		ctx := context.TODO()

		mBind := mocks.FileBinder{Mock: mock.Mock{}}
		mBind.EXPECT().CheckFileExist(ctx, "username").Return(false, nil)
		mBind.EXPECT().SaveStorage(ctx, models.StorageData{User: models.User{Username: "username"}}).Return(nil)

		service := fileStorage{bind: &mBind}
		err := service.CreateUser(ctx, models.User{Username: "username"})

		require.Equal(t, err, nil)
	})

	t.Run("should return error when file exists", func(t *testing.T) {
		ctx := context.TODO()

		mBind := mocks.FileBinder{Mock: mock.Mock{}}
		mBind.EXPECT().CheckFileExist(ctx, "username").Return(true, nil)

		service := fileStorage{bind: &mBind}
		err := service.CreateUser(ctx, models.User{Username: "username"})

		require.Equal(t, err, storage.ErrLoginAlreadyExists)
	})
}
