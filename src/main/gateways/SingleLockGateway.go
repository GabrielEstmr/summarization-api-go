package main_gateways

type SingleLockGateway interface {
	Lock() error
	Extend() (bool, error)
	Unlock() (bool, error)
}
