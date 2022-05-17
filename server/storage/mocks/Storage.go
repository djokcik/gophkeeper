// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"
	models "github.com/djokcik/gophkeeper/models"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

type Storage_Expecter struct {
	mock *mock.Mock
}

func (_m *Storage) EXPECT() *Storage_Expecter {
	return &Storage_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Storage) CreateUser(ctx context.Context, user models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type Storage_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//  - ctx context.Context
//  - user models.User
func (_e *Storage_Expecter) CreateUser(ctx interface{}, user interface{}) *Storage_CreateUser_Call {
	return &Storage_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *Storage_CreateUser_Call) Run(run func(ctx context.Context, user models.User)) *Storage_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.User))
	})
	return _c
}

func (_c *Storage_CreateUser_Call) Return(_a0 error) *Storage_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

// Read provides a mock function with given fields: ctx, username
func (_m *Storage) Read(ctx context.Context, username string) (models.StorageData, error) {
	ret := _m.Called(ctx, username)

	var r0 models.StorageData
	if rf, ok := ret.Get(0).(func(context.Context, string) models.StorageData); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(models.StorageData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type Storage_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//  - ctx context.Context
//  - username string
func (_e *Storage_Expecter) Read(ctx interface{}, username interface{}) *Storage_Read_Call {
	return &Storage_Read_Call{Call: _e.mock.On("Read", ctx, username)}
}

func (_c *Storage_Read_Call) Run(run func(ctx context.Context, username string)) *Storage_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storage_Read_Call) Return(_a0 models.StorageData, _a1 error) *Storage_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Save provides a mock function with given fields: ctx, data
func (_m *Storage) Save(ctx context.Context, data models.StorageData) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.StorageData) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Storage_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//  - ctx context.Context
//  - data models.StorageData
func (_e *Storage_Expecter) Save(ctx interface{}, data interface{}) *Storage_Save_Call {
	return &Storage_Save_Call{Call: _e.mock.On("Save", ctx, data)}
}

func (_c *Storage_Save_Call) Run(run func(ctx context.Context, data models.StorageData)) *Storage_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.StorageData))
	})
	return _c
}

func (_c *Storage_Save_Call) Return(_a0 error) *Storage_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

// UserByUsername provides a mock function with given fields: ctx, username
func (_m *Storage) UserByUsername(ctx context.Context, username string) (models.User, error) {
	ret := _m.Called(ctx, username)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_UserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserByUsername'
type Storage_UserByUsername_Call struct {
	*mock.Call
}

// UserByUsername is a helper method to define mock.On call
//  - ctx context.Context
//  - username string
func (_e *Storage_Expecter) UserByUsername(ctx interface{}, username interface{}) *Storage_UserByUsername_Call {
	return &Storage_UserByUsername_Call{Call: _e.mock.On("UserByUsername", ctx, username)}
}

func (_c *Storage_UserByUsername_Call) Run(run func(ctx context.Context, username string)) *Storage_UserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Storage_UserByUsername_Call) Return(_a0 models.User, _a1 error) *Storage_UserByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NewStorage creates a new instance of Storage. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t testing.TB) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
