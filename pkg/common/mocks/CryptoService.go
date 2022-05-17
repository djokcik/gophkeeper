// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// CryptoService is an autogenerated mock type for the CryptoService type
type CryptoService struct {
	mock.Mock
}

type CryptoService_Expecter struct {
	mock *mock.Mock
}

func (_m *CryptoService) EXPECT() *CryptoService_Expecter {
	return &CryptoService_Expecter{mock: &_m.Mock}
}

// Decrypt provides a mock function with given fields: ctx, encryptedString, key
func (_m *CryptoService) Decrypt(ctx context.Context, encryptedString string, key string) ([]byte, error) {
	ret := _m.Called(ctx, encryptedString, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []byte); ok {
		r0 = rf(ctx, encryptedString, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, encryptedString, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CryptoService_Decrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrypt'
type CryptoService_Decrypt_Call struct {
	*mock.Call
}

// Decrypt is a helper method to define mock.On call
//  - ctx context.Context
//  - encryptedString string
//  - key string
func (_e *CryptoService_Expecter) Decrypt(ctx interface{}, encryptedString interface{}, key interface{}) *CryptoService_Decrypt_Call {
	return &CryptoService_Decrypt_Call{Call: _e.mock.On("Decrypt", ctx, encryptedString, key)}
}

func (_c *CryptoService_Decrypt_Call) Run(run func(ctx context.Context, encryptedString string, key string)) *CryptoService_Decrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *CryptoService_Decrypt_Call) Return(_a0 []byte, _a1 error) *CryptoService_Decrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DecryptData provides a mock function with given fields: ctx, userPassword, encryptedData, response
func (_m *CryptoService) DecryptData(ctx context.Context, userPassword string, encryptedData string, response interface{}) error {
	ret := _m.Called(ctx, userPassword, encryptedData, response)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, userPassword, encryptedData, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CryptoService_DecryptData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecryptData'
type CryptoService_DecryptData_Call struct {
	*mock.Call
}

// DecryptData is a helper method to define mock.On call
//  - ctx context.Context
//  - userPassword string
//  - encryptedData string
//  - response interface{}
func (_e *CryptoService_Expecter) DecryptData(ctx interface{}, userPassword interface{}, encryptedData interface{}, response interface{}) *CryptoService_DecryptData_Call {
	return &CryptoService_DecryptData_Call{Call: _e.mock.On("DecryptData", ctx, userPassword, encryptedData, response)}
}

func (_c *CryptoService_DecryptData_Call) Run(run func(ctx context.Context, userPassword string, encryptedData string, response interface{})) *CryptoService_DecryptData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(interface{}))
	})
	return _c
}

func (_c *CryptoService_DecryptData_Call) Return(_a0 error) *CryptoService_DecryptData_Call {
	_c.Call.Return(_a0)
	return _c
}

// Encrypt provides a mock function with given fields: ctx, data, key
func (_m *CryptoService) Encrypt(ctx context.Context, data []byte, key string) (string, error) {
	ret := _m.Called(ctx, data, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) string); ok {
		r0 = rf(ctx, data, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, data, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CryptoService_Encrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encrypt'
type CryptoService_Encrypt_Call struct {
	*mock.Call
}

// Encrypt is a helper method to define mock.On call
//  - ctx context.Context
//  - data []byte
//  - key string
func (_e *CryptoService_Expecter) Encrypt(ctx interface{}, data interface{}, key interface{}) *CryptoService_Encrypt_Call {
	return &CryptoService_Encrypt_Call{Call: _e.mock.On("Encrypt", ctx, data, key)}
}

func (_c *CryptoService_Encrypt_Call) Run(run func(ctx context.Context, data []byte, key string)) *CryptoService_Encrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(string))
	})
	return _c
}

func (_c *CryptoService_Encrypt_Call) Return(_a0 string, _a1 error) *CryptoService_Encrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// EncryptData provides a mock function with given fields: ctx, userPassword, data
func (_m *CryptoService) EncryptData(ctx context.Context, userPassword string, data interface{}) (string, error) {
	ret := _m.Called(ctx, userPassword, data)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) string); ok {
		r0 = rf(ctx, userPassword, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, userPassword, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CryptoService_EncryptData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EncryptData'
type CryptoService_EncryptData_Call struct {
	*mock.Call
}

// EncryptData is a helper method to define mock.On call
//  - ctx context.Context
//  - userPassword string
//  - data interface{}
func (_e *CryptoService_Expecter) EncryptData(ctx interface{}, userPassword interface{}, data interface{}) *CryptoService_EncryptData_Call {
	return &CryptoService_EncryptData_Call{Call: _e.mock.On("EncryptData", ctx, userPassword, data)}
}

func (_c *CryptoService_EncryptData_Call) Run(run func(ctx context.Context, userPassword string, data interface{})) *CryptoService_EncryptData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *CryptoService_EncryptData_Call) Return(_a0 string, _a1 error) *CryptoService_EncryptData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GenerateHash provides a mock function with given fields: value
func (_m *CryptoService) GenerateHash(value string) string {
	ret := _m.Called(value)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CryptoService_GenerateHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateHash'
type CryptoService_GenerateHash_Call struct {
	*mock.Call
}

// GenerateHash is a helper method to define mock.On call
//  - value string
func (_e *CryptoService_Expecter) GenerateHash(value interface{}) *CryptoService_GenerateHash_Call {
	return &CryptoService_GenerateHash_Call{Call: _e.mock.On("GenerateHash", value)}
}

func (_c *CryptoService_GenerateHash_Call) Run(run func(value string)) *CryptoService_GenerateHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *CryptoService_GenerateHash_Call) Return(_a0 string) *CryptoService_GenerateHash_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewCryptoService creates a new instance of CryptoService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCryptoService(t testing.TB) *CryptoService {
	mock := &CryptoService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
