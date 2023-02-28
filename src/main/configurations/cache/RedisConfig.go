package main_configurations_cache

import (
	"context"
	"log"
	main_configurations_yml "mpindicator/main/configurations/yml"
	"sync"

	"github.com/go-redis/redis/v8"
)

const MSG_ATTEMPT_TO_CONNECT_TO_REDIS_CLIENT = "Attempt to connect to redis client %s"
const MSG_SUCECSSFULLY_CONNECTED_TO_REDIS_CLIENT = "Successfully connected to redis client %s"

const REDIS_HOST_NODE1 = "Redis.hosts"

var RedisCluster *redis.Client = nil
var once sync.Once

func GetRedisClusterBean() *redis.Client {
	once.Do(func() {

		if RedisCluster == nil {
			RedisCluster = getRedisCluster()
		}

	})
	return RedisCluster
}

func getRedisCluster() *redis.Client {
	redisHost := main_configurations_yml.GetBeanPropertyByName(REDIS_HOST_NODE1)

	log.Printf(MSG_ATTEMPT_TO_CONNECT_TO_REDIS_CLIENT, redisHost)

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping(context.TODO()).Err()
	if err != nil {
		panic(err)
	}

	log.Printf(MSG_SUCECSSFULLY_CONNECTED_TO_REDIS_CLIENT, redisHost)
	return client
}
