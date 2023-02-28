package main_configurations_distributedlock

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	main_configurations_cache "mpindicator/main/configurations/cache"
	"sync"
)

var RedisCluster *redsync.Redsync = nil
var once sync.Once

func GetLockClientBean() *redsync.Redsync {
	once.Do(func() {

		if RedisCluster == nil {
			RedisCluster = getLockClient()
		}

	})
	return RedisCluster
}

func getLockClient() *redsync.Redsync {
	redisClient := main_configurations_cache.GetRedisClusterBean()
	pool := goredis.NewPool(redisClient)
	rs := redsync.New(pool)
	return rs
}
