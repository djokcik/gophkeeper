// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ServerRecordBinaryDataService is an autogenerated mock type for the ServerRecordBinaryDataService type
type ServerRecordBinaryDataService struct {
	mock.Mock
}

// Load provides a mock function with given fields: ctx, key, username
func (_m *ServerRecordBinaryDataService) Load(ctx context.Context, key string, username string) (string, error) {
	ret := _m.Called(ctx, key, username)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, key, username)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, key, username
func (_m *ServerRecordBinaryDataService) Remove(ctx context.Context, key string, username string) error {
	ret := _m.Called(ctx, key, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, key, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, key, username, data
func (_m *ServerRecordBinaryDataService) Save(ctx context.Context, key string, username string, data string) error {
	ret := _m.Called(ctx, key, username, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, key, username, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServerRecordBinaryDataService creates a new instance of ServerRecordBinaryDataService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewServerRecordBinaryDataService(t testing.TB) *ServerRecordBinaryDataService {
	mock := &ServerRecordBinaryDataService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
