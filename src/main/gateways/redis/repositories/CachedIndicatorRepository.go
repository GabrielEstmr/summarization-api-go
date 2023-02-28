package main_gateways_redis_repositories

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	main_configurations_cache "mpindicator/main/configurations/cache"
	main_gateways_redis_documents "mpindicator/main/gateways/redis/documents"
	"time"
)

type CachedIndicatorRepository struct {
	redisClient *redis.Client
}

func NewCachedUserRepository() *CachedIndicatorRepository {
	return &CachedIndicatorRepository{redisClient: main_configurations_cache.GetRedisClusterBean()}
}

func (thisRepository *CachedIndicatorRepository) FindById(indicatorId string) (
	main_gateways_redis_documents.CachedIndicatorDocument, error) {

	result, err := thisRepository.redisClient.Get(context.TODO(), indicatorId).Result()

	if err == redis.Nil {
		return main_gateways_redis_documents.CachedIndicatorDocument{}, nil
	}

	if err != nil {
		return main_gateways_redis_documents.CachedIndicatorDocument{}, err
	}

	var cachedIndicatorDocument main_gateways_redis_documents.CachedIndicatorDocument
	if err = json.Unmarshal([]byte(result), &cachedIndicatorDocument); err != nil {
		return main_gateways_redis_documents.CachedIndicatorDocument{}, err
	}

	return cachedIndicatorDocument, nil
}

func (thisRepository *CachedIndicatorRepository) Save(
	userDocument main_gateways_redis_documents.CachedIndicatorDocument,
) (main_gateways_redis_documents.CachedIndicatorDocument, error) {

	userBytes, err := json.Marshal(userDocument)
	if err != nil {
		return main_gateways_redis_documents.CachedIndicatorDocument{}, err
	}

	val, err := thisRepository.redisClient.Set(context.TODO(), userDocument.Id, userBytes, time.Hour).Result()
	log.Println(val)

	if err != nil {
		return main_gateways_redis_documents.CachedIndicatorDocument{}, err
	}

	return userDocument, nil
}
