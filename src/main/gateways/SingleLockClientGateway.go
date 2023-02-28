package main_gateways

import "time"

type SingleLockClientGateway interface {
	GetClient(key string, timeToLive time.Duration)
}
