package service

import "sync"

//go:generate mockery --name=KeyLockService

// KeyLockService provides interface for lock by username
type KeyLockService interface {
	GetLockBy(key string) *sync.Mutex
	Lock(key string)
	Unlock(key string)
}

type stringKeyLock struct {
	mapLock sync.Mutex // to make the map safe concurrently
	locks   map[string]*sync.Mutex
}

func NewStringKeyLock() *stringKeyLock {
	return &stringKeyLock{locks: make(map[string]*sync.Mutex)}
}

// GetLockBy returns lock by name
func (l *stringKeyLock) GetLockBy(key string) *sync.Mutex {
	l.mapLock.Lock()
	defer l.mapLock.Unlock()

	ret, found := l.locks[key]
	if found {
		return ret
	}

	ret = &sync.Mutex{}
	l.locks[key] = ret
	return ret
}

// Lock is lock by name
func (l *stringKeyLock) Lock(key string) {
	l.GetLockBy(key).Lock()
}

// Unlock is unlock by name
func (l *stringKeyLock) Unlock(key string) {
	l.GetLockBy(key).Unlock()
}
