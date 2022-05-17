// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ClientRpcService is an autogenerated mock type for the ClientRpcService type
type ClientRpcService struct {
	mock.Mock
}

type ClientRpcService_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientRpcService) EXPECT() *ClientRpcService_Expecter {
	return &ClientRpcService_Expecter{mock: &_m.Mock}
}

// Call provides a mock function with given fields: ctx, serviceMethod, args, reply
func (_m *ClientRpcService) Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	ret := _m.Called(ctx, serviceMethod, args, reply)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, interface{}) error); ok {
		r0 = rf(ctx, serviceMethod, args, reply)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_Call_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Call'
type ClientRpcService_Call_Call struct {
	*mock.Call
}

// Call is a helper method to define mock.On call
//  - ctx context.Context
//  - serviceMethod string
//  - args interface{}
//  - reply interface{}
func (_e *ClientRpcService_Expecter) Call(ctx interface{}, serviceMethod interface{}, args interface{}, reply interface{}) *ClientRpcService_Call_Call {
	return &ClientRpcService_Call_Call{Call: _e.mock.On("Call", ctx, serviceMethod, args, reply)}
}

func (_c *ClientRpcService_Call_Call) Run(run func(ctx context.Context, serviceMethod string, args interface{}, reply interface{})) *ClientRpcService_Call_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(interface{}))
	})
	return _c
}

func (_c *ClientRpcService_Call_Call) Return(_a0 error) *ClientRpcService_Call_Call {
	_c.Call.Return(_a0)
	return _c
}

// CheckOnline provides a mock function with given fields:
func (_m *ClientRpcService) CheckOnline() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ClientRpcService_CheckOnline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckOnline'
type ClientRpcService_CheckOnline_Call struct {
	*mock.Call
}

// CheckOnline is a helper method to define mock.On call
func (_e *ClientRpcService_Expecter) CheckOnline() *ClientRpcService_CheckOnline_Call {
	return &ClientRpcService_CheckOnline_Call{Call: _e.mock.On("CheckOnline")}
}

func (_c *ClientRpcService_CheckOnline_Call) Run(run func()) *ClientRpcService_CheckOnline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClientRpcService_CheckOnline_Call) Return(_a0 bool) *ClientRpcService_CheckOnline_Call {
	_c.Call.Return(_a0)
	return _c
}

// LoadRecordBankCardByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) LoadRecordBankCardByKey(ctx context.Context, token string, key string) (string, error) {
	ret := _m.Called(ctx, token, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, token, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRpcService_LoadRecordBankCardByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadRecordBankCardByKey'
type ClientRpcService_LoadRecordBankCardByKey_Call struct {
	*mock.Call
}

// LoadRecordBankCardByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) LoadRecordBankCardByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_LoadRecordBankCardByKey_Call {
	return &ClientRpcService_LoadRecordBankCardByKey_Call{Call: _e.mock.On("LoadRecordBankCardByKey", ctx, token, key)}
}

func (_c *ClientRpcService_LoadRecordBankCardByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_LoadRecordBankCardByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_LoadRecordBankCardByKey_Call) Return(_a0 string, _a1 error) *ClientRpcService_LoadRecordBankCardByKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadRecordBinaryDataByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) LoadRecordBinaryDataByKey(ctx context.Context, token string, key string) (string, error) {
	ret := _m.Called(ctx, token, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, token, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRpcService_LoadRecordBinaryDataByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadRecordBinaryDataByKey'
type ClientRpcService_LoadRecordBinaryDataByKey_Call struct {
	*mock.Call
}

// LoadRecordBinaryDataByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) LoadRecordBinaryDataByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_LoadRecordBinaryDataByKey_Call {
	return &ClientRpcService_LoadRecordBinaryDataByKey_Call{Call: _e.mock.On("LoadRecordBinaryDataByKey", ctx, token, key)}
}

func (_c *ClientRpcService_LoadRecordBinaryDataByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_LoadRecordBinaryDataByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_LoadRecordBinaryDataByKey_Call) Return(_a0 string, _a1 error) *ClientRpcService_LoadRecordBinaryDataByKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadRecordPersonalDataByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) LoadRecordPersonalDataByKey(ctx context.Context, token string, key string) (string, error) {
	ret := _m.Called(ctx, token, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, token, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRpcService_LoadRecordPersonalDataByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadRecordPersonalDataByKey'
type ClientRpcService_LoadRecordPersonalDataByKey_Call struct {
	*mock.Call
}

// LoadRecordPersonalDataByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) LoadRecordPersonalDataByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_LoadRecordPersonalDataByKey_Call {
	return &ClientRpcService_LoadRecordPersonalDataByKey_Call{Call: _e.mock.On("LoadRecordPersonalDataByKey", ctx, token, key)}
}

func (_c *ClientRpcService_LoadRecordPersonalDataByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_LoadRecordPersonalDataByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_LoadRecordPersonalDataByKey_Call) Return(_a0 string, _a1 error) *ClientRpcService_LoadRecordPersonalDataByKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadRecordTextDataByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) LoadRecordTextDataByKey(ctx context.Context, token string, key string) (string, error) {
	ret := _m.Called(ctx, token, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, token, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRpcService_LoadRecordTextDataByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadRecordTextDataByKey'
type ClientRpcService_LoadRecordTextDataByKey_Call struct {
	*mock.Call
}

// LoadRecordTextDataByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) LoadRecordTextDataByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_LoadRecordTextDataByKey_Call {
	return &ClientRpcService_LoadRecordTextDataByKey_Call{Call: _e.mock.On("LoadRecordTextDataByKey", ctx, token, key)}
}

func (_c *ClientRpcService_LoadRecordTextDataByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_LoadRecordTextDataByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_LoadRecordTextDataByKey_Call) Return(_a0 string, _a1 error) *ClientRpcService_LoadRecordTextDataByKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Login provides a mock function with given fields: ctx, username, password
func (_m *ClientRpcService) Login(ctx context.Context, username string, password string) (string, error) {
	ret := _m.Called(ctx, username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRpcService_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type ClientRpcService_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//  - ctx context.Context
//  - username string
//  - password string
func (_e *ClientRpcService_Expecter) Login(ctx interface{}, username interface{}, password interface{}) *ClientRpcService_Login_Call {
	return &ClientRpcService_Login_Call{Call: _e.mock.On("Login", ctx, username, password)}
}

func (_c *ClientRpcService_Login_Call) Run(run func(ctx context.Context, username string, password string)) *ClientRpcService_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_Login_Call) Return(_a0 string, _a1 error) *ClientRpcService_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Register provides a mock function with given fields: ctx, username, password
func (_m *ClientRpcService) Register(ctx context.Context, username string, password string) (string, error) {
	ret := _m.Called(ctx, username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientRpcService_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type ClientRpcService_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//  - ctx context.Context
//  - username string
//  - password string
func (_e *ClientRpcService_Expecter) Register(ctx interface{}, username interface{}, password interface{}) *ClientRpcService_Register_Call {
	return &ClientRpcService_Register_Call{Call: _e.mock.On("Register", ctx, username, password)}
}

func (_c *ClientRpcService_Register_Call) Run(run func(ctx context.Context, username string, password string)) *ClientRpcService_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_Register_Call) Return(_a0 string, _a1 error) *ClientRpcService_Register_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RemoveRecordBankCardByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) RemoveRecordBankCardByKey(ctx context.Context, token string, key string) error {
	ret := _m.Called(ctx, token, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_RemoveRecordBankCardByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRecordBankCardByKey'
type ClientRpcService_RemoveRecordBankCardByKey_Call struct {
	*mock.Call
}

// RemoveRecordBankCardByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) RemoveRecordBankCardByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_RemoveRecordBankCardByKey_Call {
	return &ClientRpcService_RemoveRecordBankCardByKey_Call{Call: _e.mock.On("RemoveRecordBankCardByKey", ctx, token, key)}
}

func (_c *ClientRpcService_RemoveRecordBankCardByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_RemoveRecordBankCardByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_RemoveRecordBankCardByKey_Call) Return(_a0 error) *ClientRpcService_RemoveRecordBankCardByKey_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveRecordBinaryDataByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) RemoveRecordBinaryDataByKey(ctx context.Context, token string, key string) error {
	ret := _m.Called(ctx, token, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_RemoveRecordBinaryDataByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRecordBinaryDataByKey'
type ClientRpcService_RemoveRecordBinaryDataByKey_Call struct {
	*mock.Call
}

// RemoveRecordBinaryDataByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) RemoveRecordBinaryDataByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_RemoveRecordBinaryDataByKey_Call {
	return &ClientRpcService_RemoveRecordBinaryDataByKey_Call{Call: _e.mock.On("RemoveRecordBinaryDataByKey", ctx, token, key)}
}

func (_c *ClientRpcService_RemoveRecordBinaryDataByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_RemoveRecordBinaryDataByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_RemoveRecordBinaryDataByKey_Call) Return(_a0 error) *ClientRpcService_RemoveRecordBinaryDataByKey_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveRecordPersonalDataByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) RemoveRecordPersonalDataByKey(ctx context.Context, token string, key string) error {
	ret := _m.Called(ctx, token, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_RemoveRecordPersonalDataByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRecordPersonalDataByKey'
type ClientRpcService_RemoveRecordPersonalDataByKey_Call struct {
	*mock.Call
}

// RemoveRecordPersonalDataByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) RemoveRecordPersonalDataByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_RemoveRecordPersonalDataByKey_Call {
	return &ClientRpcService_RemoveRecordPersonalDataByKey_Call{Call: _e.mock.On("RemoveRecordPersonalDataByKey", ctx, token, key)}
}

func (_c *ClientRpcService_RemoveRecordPersonalDataByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_RemoveRecordPersonalDataByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_RemoveRecordPersonalDataByKey_Call) Return(_a0 error) *ClientRpcService_RemoveRecordPersonalDataByKey_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveRecordTextDataByKey provides a mock function with given fields: ctx, token, key
func (_m *ClientRpcService) RemoveRecordTextDataByKey(ctx context.Context, token string, key string) error {
	ret := _m.Called(ctx, token, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_RemoveRecordTextDataByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRecordTextDataByKey'
type ClientRpcService_RemoveRecordTextDataByKey_Call struct {
	*mock.Call
}

// RemoveRecordTextDataByKey is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
func (_e *ClientRpcService_Expecter) RemoveRecordTextDataByKey(ctx interface{}, token interface{}, key interface{}) *ClientRpcService_RemoveRecordTextDataByKey_Call {
	return &ClientRpcService_RemoveRecordTextDataByKey_Call{Call: _e.mock.On("RemoveRecordTextDataByKey", ctx, token, key)}
}

func (_c *ClientRpcService_RemoveRecordTextDataByKey_Call) Run(run func(ctx context.Context, token string, key string)) *ClientRpcService_RemoveRecordTextDataByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClientRpcService_RemoveRecordTextDataByKey_Call) Return(_a0 error) *ClientRpcService_RemoveRecordTextDataByKey_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveRecordBankCard provides a mock function with given fields: ctx, token, key, data
func (_m *ClientRpcService) SaveRecordBankCard(ctx context.Context, token string, key string, data string) error {
	ret := _m.Called(ctx, token, key, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, token, key, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_SaveRecordBankCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveRecordBankCard'
type ClientRpcService_SaveRecordBankCard_Call struct {
	*mock.Call
}

// SaveRecordBankCard is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
//  - data string
func (_e *ClientRpcService_Expecter) SaveRecordBankCard(ctx interface{}, token interface{}, key interface{}, data interface{}) *ClientRpcService_SaveRecordBankCard_Call {
	return &ClientRpcService_SaveRecordBankCard_Call{Call: _e.mock.On("SaveRecordBankCard", ctx, token, key, data)}
}

func (_c *ClientRpcService_SaveRecordBankCard_Call) Run(run func(ctx context.Context, token string, key string, data string)) *ClientRpcService_SaveRecordBankCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ClientRpcService_SaveRecordBankCard_Call) Return(_a0 error) *ClientRpcService_SaveRecordBankCard_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveRecordBinaryData provides a mock function with given fields: ctx, token, key, data
func (_m *ClientRpcService) SaveRecordBinaryData(ctx context.Context, token string, key string, data string) error {
	ret := _m.Called(ctx, token, key, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, token, key, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_SaveRecordBinaryData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveRecordBinaryData'
type ClientRpcService_SaveRecordBinaryData_Call struct {
	*mock.Call
}

// SaveRecordBinaryData is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
//  - data string
func (_e *ClientRpcService_Expecter) SaveRecordBinaryData(ctx interface{}, token interface{}, key interface{}, data interface{}) *ClientRpcService_SaveRecordBinaryData_Call {
	return &ClientRpcService_SaveRecordBinaryData_Call{Call: _e.mock.On("SaveRecordBinaryData", ctx, token, key, data)}
}

func (_c *ClientRpcService_SaveRecordBinaryData_Call) Run(run func(ctx context.Context, token string, key string, data string)) *ClientRpcService_SaveRecordBinaryData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ClientRpcService_SaveRecordBinaryData_Call) Return(_a0 error) *ClientRpcService_SaveRecordBinaryData_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveRecordPersonalData provides a mock function with given fields: ctx, token, key, data
func (_m *ClientRpcService) SaveRecordPersonalData(ctx context.Context, token string, key string, data string) error {
	ret := _m.Called(ctx, token, key, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, token, key, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_SaveRecordPersonalData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveRecordPersonalData'
type ClientRpcService_SaveRecordPersonalData_Call struct {
	*mock.Call
}

// SaveRecordPersonalData is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
//  - data string
func (_e *ClientRpcService_Expecter) SaveRecordPersonalData(ctx interface{}, token interface{}, key interface{}, data interface{}) *ClientRpcService_SaveRecordPersonalData_Call {
	return &ClientRpcService_SaveRecordPersonalData_Call{Call: _e.mock.On("SaveRecordPersonalData", ctx, token, key, data)}
}

func (_c *ClientRpcService_SaveRecordPersonalData_Call) Run(run func(ctx context.Context, token string, key string, data string)) *ClientRpcService_SaveRecordPersonalData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ClientRpcService_SaveRecordPersonalData_Call) Return(_a0 error) *ClientRpcService_SaveRecordPersonalData_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveRecordTextData provides a mock function with given fields: ctx, token, key, data
func (_m *ClientRpcService) SaveRecordTextData(ctx context.Context, token string, key string, data string) error {
	ret := _m.Called(ctx, token, key, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, token, key, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientRpcService_SaveRecordTextData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveRecordTextData'
type ClientRpcService_SaveRecordTextData_Call struct {
	*mock.Call
}

// SaveRecordTextData is a helper method to define mock.On call
//  - ctx context.Context
//  - token string
//  - key string
//  - data string
func (_e *ClientRpcService_Expecter) SaveRecordTextData(ctx interface{}, token interface{}, key interface{}, data interface{}) *ClientRpcService_SaveRecordTextData_Call {
	return &ClientRpcService_SaveRecordTextData_Call{Call: _e.mock.On("SaveRecordTextData", ctx, token, key, data)}
}

func (_c *ClientRpcService_SaveRecordTextData_Call) Run(run func(ctx context.Context, token string, key string, data string)) *ClientRpcService_SaveRecordTextData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ClientRpcService_SaveRecordTextData_Call) Return(_a0 error) *ClientRpcService_SaveRecordTextData_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewClientRpcService creates a new instance of ClientRpcService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientRpcService(t testing.TB) *ClientRpcService {
	mock := &ClientRpcService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
