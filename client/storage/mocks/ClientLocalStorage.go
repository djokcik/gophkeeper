// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"
	clientmodels "github.com/djokcik/gophkeeper/client/clientmodels"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ClientLocalStorage is an autogenerated mock type for the ClientLocalStorage type
type ClientLocalStorage struct {
	mock.Mock
}

type ClientLocalStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientLocalStorage) EXPECT() *ClientLocalStorage_Expecter {
	return &ClientLocalStorage_Expecter{mock: &_m.Mock}
}

// ClearActions provides a mock function with given fields: ctx
func (_m *ClientLocalStorage) ClearActions(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientLocalStorage_ClearActions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearActions'
type ClientLocalStorage_ClearActions_Call struct {
	*mock.Call
}

// ClearActions is a helper method to define mock.On call
//  - ctx context.Context
func (_e *ClientLocalStorage_Expecter) ClearActions(ctx interface{}) *ClientLocalStorage_ClearActions_Call {
	return &ClientLocalStorage_ClearActions_Call{Call: _e.mock.On("ClearActions", ctx)}
}

func (_c *ClientLocalStorage_ClearActions_Call) Run(run func(ctx context.Context)) *ClientLocalStorage_ClearActions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ClientLocalStorage_ClearActions_Call) Return(_a0 error) *ClientLocalStorage_ClearActions_Call {
	_c.Call.Return(_a0)
	return _c
}

// Close provides a mock function with given fields:
func (_m *ClientLocalStorage) Close() {
	_m.Called()
}

// ClientLocalStorage_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type ClientLocalStorage_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *ClientLocalStorage_Expecter) Close() *ClientLocalStorage_Close_Call {
	return &ClientLocalStorage_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *ClientLocalStorage_Close_Call) Run(run func()) *ClientLocalStorage_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClientLocalStorage_Close_Call) Return() *ClientLocalStorage_Close_Call {
	_c.Call.Return()
	return _c
}

// LoadRecords provides a mock function with given fields: ctx
func (_m *ClientLocalStorage) LoadRecords(ctx context.Context) ([]clientmodels.RecordFileLine, error) {
	ret := _m.Called(ctx)

	var r0 []clientmodels.RecordFileLine
	if rf, ok := ret.Get(0).(func(context.Context) []clientmodels.RecordFileLine); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clientmodels.RecordFileLine)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientLocalStorage_LoadRecords_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadRecords'
type ClientLocalStorage_LoadRecords_Call struct {
	*mock.Call
}

// LoadRecords is a helper method to define mock.On call
//  - ctx context.Context
func (_e *ClientLocalStorage_Expecter) LoadRecords(ctx interface{}) *ClientLocalStorage_LoadRecords_Call {
	return &ClientLocalStorage_LoadRecords_Call{Call: _e.mock.On("LoadRecords", ctx)}
}

func (_c *ClientLocalStorage_LoadRecords_Call) Run(run func(ctx context.Context)) *ClientLocalStorage_LoadRecords_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ClientLocalStorage_LoadRecords_Call) Return(_a0 []clientmodels.RecordFileLine, _a1 error) *ClientLocalStorage_LoadRecords_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RemoveRecord provides a mock function with given fields: ctx, key, method
func (_m *ClientLocalStorage) RemoveRecord(ctx context.Context, key string, method string) error {
	ret := _m.Called(ctx, key, method)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, key, method)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientLocalStorage_RemoveRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRecord'
type ClientLocalStorage_RemoveRecord_Call struct {
	*mock.Call
}

// RemoveRecord is a helper method to define mock.On call
//  - ctx context.Context
//  - key string
//  - method string
func (_e *ClientLocalStorage_Expecter) RemoveRecord(ctx interface{}, key interface{}, method interface{}) *ClientLocalStorage_RemoveRecord_Call {
	return &ClientLocalStorage_RemoveRecord_Call{Call: _e.mock.On("RemoveRecord", ctx, key, method)}
}

func (_c *ClientLocalStorage_RemoveRecord_Call) Run(run func(ctx context.Context, key string, method string)) *ClientLocalStorage_RemoveRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientLocalStorage_RemoveRecord_Call) Return(_a0 error) *ClientLocalStorage_RemoveRecord_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveRecord provides a mock function with given fields: ctx, key, data, method
func (_m *ClientLocalStorage) SaveRecord(ctx context.Context, key string, data string, method string) error {
	ret := _m.Called(ctx, key, data, method)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, key, data, method)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientLocalStorage_SaveRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveRecord'
type ClientLocalStorage_SaveRecord_Call struct {
	*mock.Call
}

// SaveRecord is a helper method to define mock.On call
//  - ctx context.Context
//  - key string
//  - data string
//  - method string
func (_e *ClientLocalStorage_Expecter) SaveRecord(ctx interface{}, key interface{}, data interface{}, method interface{}) *ClientLocalStorage_SaveRecord_Call {
	return &ClientLocalStorage_SaveRecord_Call{Call: _e.mock.On("SaveRecord", ctx, key, data, method)}
}

func (_c *ClientLocalStorage_SaveRecord_Call) Run(run func(ctx context.Context, key string, data string, method string)) *ClientLocalStorage_SaveRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ClientLocalStorage_SaveRecord_Call) Return(_a0 error) *ClientLocalStorage_SaveRecord_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewClientLocalStorage creates a new instance of ClientLocalStorage. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientLocalStorage(t testing.TB) *ClientLocalStorage {
	mock := &ClientLocalStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
