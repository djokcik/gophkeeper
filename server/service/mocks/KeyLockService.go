// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sync "sync"

	testing "testing"
)

// KeyLockService is an autogenerated mock type for the KeyLockService type
type KeyLockService struct {
	mock.Mock
}

// GetLockBy provides a mock function with given fields: key
func (_m *KeyLockService) GetLockBy(key string) *sync.Mutex {
	ret := _m.Called(key)

	var r0 *sync.Mutex
	if rf, ok := ret.Get(0).(func(string) *sync.Mutex); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sync.Mutex)
		}
	}

	return r0
}

// Lock provides a mock function with given fields: key
func (_m *KeyLockService) Lock(key string) {
	_m.Called(key)
}

// Unlock provides a mock function with given fields: key
func (_m *KeyLockService) Unlock(key string) {
	_m.Called(key)
}

// NewKeyLockService creates a new instance of KeyLockService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewKeyLockService(t testing.TB) *KeyLockService {
	mock := &KeyLockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}