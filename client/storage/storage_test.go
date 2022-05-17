package storage

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/storage/mocks"
	"testing"
)

func Test_localStorage_ClearActions(t *testing.T) {
	t.Run("actions should be cleared", func(t *testing.T) {
		mFileWriter := mocks.ActionFileStoreWriter{Mock: mock.Mock{}}
		mFileWriter.EXPECT().SaveActions(clientmodels.StoreActions{}).Return(nil)

		svc := localStorage{FileWriter: &mFileWriter}
		err := svc.ClearActions(context.TODO())

		require.Equal(t, err, nil)
	})
}

func Test_localStorage_LoadRecords(t *testing.T) {
	t.Run("actions should be loaded", func(t *testing.T) {
		mFileReader := mocks.ActionFileStoreReader{Mock: mock.Mock{}}
		mFileReader.EXPECT().ReadActions().Return([]clientmodels.RecordFileLine{{Key: "testKey"}}, nil)

		svc := localStorage{FileReader: &mFileReader}
		lines, err := svc.LoadRecords(context.TODO())

		require.Equal(t, err, nil)
		require.Equal(t, lines, []clientmodels.RecordFileLine{{Key: "testKey"}})
	})
}

func Test_localStorage_RemoveRecord(t *testing.T) {
	t.Run("remove action should be added", func(t *testing.T) {
		ctx := context.TODO()

		mFileReader := mocks.ActionFileStoreReader{Mock: mock.Mock{}}
		mFileReader.EXPECT().ReadActions().Return([]clientmodels.RecordFileLine{{Key: "testKey1"}}, nil)

		mFileWriter := mocks.ActionFileStoreWriter{Mock: mock.Mock{}}
		mFileWriter.EXPECT().SaveActions(clientmodels.StoreActions{
			{Key: "testKey1"},
			{Key: "testKey2", Method: "testActionMethod", ActionType: clientmodels.RemoveMethod},
		}).Return(nil)

		svc := localStorage{FileReader: &mFileReader, FileWriter: &mFileWriter}
		err := svc.RemoveRecord(ctx, "testKey2", "testActionMethod")

		require.Equal(t, err, nil)
	})
}

func Test_localStorage_SaveRecord(t *testing.T) {
	t.Run("save action should be added", func(t *testing.T) {
		ctx := context.TODO()

		mFileReader := mocks.ActionFileStoreReader{Mock: mock.Mock{}}
		mFileReader.EXPECT().ReadActions().Return([]clientmodels.RecordFileLine{{Key: "testKey1"}}, nil)

		mFileWriter := mocks.ActionFileStoreWriter{Mock: mock.Mock{}}
		mFileWriter.EXPECT().SaveActions(clientmodels.StoreActions{
			{Key: "testKey1"},
			{Key: "testKey2", Method: "testActionMethod", ActionType: clientmodels.SaveMethod, Data: "testData"},
		}).Return(nil)

		svc := localStorage{FileReader: &mFileReader, FileWriter: &mFileWriter}
		err := svc.SaveRecord(ctx, "testKey2", "testData", "testActionMethod")

		require.Equal(t, err, nil)
	})
}

func Test_localStorage_Close(t *testing.T) {
	t.Run("files should be closed", func(t *testing.T) {
		mFileReader := mocks.ActionFileStoreReader{Mock: mock.Mock{}}
		mFileReader.EXPECT().Close().Return(nil)

		mFileWriter := mocks.ActionFileStoreWriter{Mock: mock.Mock{}}
		mFileWriter.EXPECT().Close().Return(nil)

		svc := localStorage{FileReader: &mFileReader, FileWriter: &mFileWriter}
		svc.Close()

		mFileReader.AssertNumberOfCalls(t, "Close", 1)
		mFileWriter.AssertNumberOfCalls(t, "Close", 1)
	})
}
