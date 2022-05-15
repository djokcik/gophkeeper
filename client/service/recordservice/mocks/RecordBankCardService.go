// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"
	clientmodels "gophkeeper/client/clientmodels"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RecordBankCardService is an autogenerated mock type for the RecordBankCardService type
type RecordBankCardService struct {
	mock.Mock
}

// LoadRecordByKey provides a mock function with given fields: ctx, key
func (_m *RecordBankCardService) LoadRecordByKey(ctx context.Context, key string) (clientmodels.RecordBankCardData, error) {
	ret := _m.Called(ctx, key)

	var r0 clientmodels.RecordBankCardData
	if rf, ok := ret.Get(0).(func(context.Context, string) clientmodels.RecordBankCardData); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(clientmodels.RecordBankCardData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRecordByKey provides a mock function with given fields: ctx, key
func (_m *RecordBankCardService) RemoveRecordByKey(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRecord provides a mock function with given fields: ctx, key, data
func (_m *RecordBankCardService) SaveRecord(ctx context.Context, key string, data clientmodels.RecordBankCardData) error {
	ret := _m.Called(ctx, key, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, clientmodels.RecordBankCardData) error); ok {
		r0 = rf(ctx, key, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRecordBankCardService creates a new instance of RecordBankCardService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRecordBankCardService(t testing.TB) *RecordBankCardService {
	mock := &RecordBankCardService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
