package service

import "sync"

type KeyLockService interface {
	getLockBy(key string) *sync.Mutex
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

func (l *stringKeyLock) getLockBy(key string) *sync.Mutex {
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

func (l *stringKeyLock) Lock(key string) {
	l.getLockBy(key).Lock()
}

func (l *stringKeyLock) Unlock(key string) {
	l.getLockBy(key).Unlock()
}
