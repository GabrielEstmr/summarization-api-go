package main_domains

import (
	main_gateways_distributedlock "mpindicator/main/gateways/distributedlock"
	"time"
)

type SingleLock struct {
	singleLockAdapter main_gateways_distributedlock.SingleLockAdapter
}

func NewSingleLock(key string, timeToLive time.Duration) *SingleLock {
	return &SingleLock{
		*main_gateways_distributedlock.NewSingleLockAdapter(key, timeToLive)}
}

func (thisAdapter *SingleLock) Lock() error {
	return thisAdapter.singleLockAdapter.Lock()
}

func (thisAdapter *SingleLock) Extend() (bool, error) {
	return thisAdapter.singleLockAdapter.Extend()
}

func (thisAdapter *SingleLock) Unlock() (bool, error) {
	return thisAdapter.singleLockAdapter.Unlock()
}
