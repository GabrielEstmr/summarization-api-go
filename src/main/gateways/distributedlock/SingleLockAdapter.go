package main_gateways_distributedlock

import (
	main_configurations_distributedlock "mpindicator/main/configurations/distributedlock"
	"time"

	"github.com/go-redsync/redsync/v4"
)

type SingleLockAdapter struct {
	lock *redsync.Mutex
}

func NewSingleLockAdapter(key string, timeToLive time.Duration) *SingleLockAdapter {
	lock := main_configurations_distributedlock.GetLockClientBean().NewMutex(key, redsync.WithExpiry(timeToLive))
	return &SingleLockAdapter{
		lock: lock,
	}
}

func (thisAdapter *SingleLockAdapter) Lock() error {
	return thisAdapter.lock.Lock()
}

func (thisAdapter *SingleLockAdapter) Extend() (bool, error) {
	return thisAdapter.lock.Extend()
}

func (thisAdapter *SingleLockAdapter) Unlock() (bool, error) {
	return thisAdapter.lock.Unlock()
}
