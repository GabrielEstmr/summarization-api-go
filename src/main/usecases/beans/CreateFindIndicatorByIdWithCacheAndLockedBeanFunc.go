package main_usecases_beans

import (
	main_gateways "mpindicator/main/gateways"
	main_gateways_redis "mpindicator/main/gateways/redis"
	main_gateways_redis_repositories "mpindicator/main/gateways/redis/repositories"

	main_gateways_mongodb "mpindicator/main/gateways/mongodb"
	main_gateways_mongodb_repositories "mpindicator/main/gateways/mongodb/repositories"
	main_usecases "mpindicator/main/usecases"
)

func CreateFindIndicatorByIdWithCacheAndLockedBeanFunc() *main_usecases.FindIndicatorByIdWithCacheAndLocked {

	var indicatorDatabaseGateway main_gateways.IndicatorDatabaseGateway
	indicatorRepository := main_gateways_mongodb_repositories.NewIndicatorRepository()
	indicatorDatabaseGatewayImpl := main_gateways_mongodb.NewIndicatorDatabaseGatewayImpl(indicatorRepository)
	indicatorDatabaseGateway = indicatorDatabaseGatewayImpl

	findIndicatorById := main_usecases.NewFindIndicatorById(&indicatorDatabaseGateway)

	var indicatorCachedDatabaseGateway main_gateways.IndicatorCachedDatabaseGateway
	cachedUserRepository := main_gateways_redis_repositories.NewCachedUserRepository()
	indicatorCachedDatabaseGatewayImpl := main_gateways_redis.NewIndicatorCachedDatabaseGatewayImpl(cachedUserRepository)
	indicatorCachedDatabaseGateway = indicatorCachedDatabaseGatewayImpl

	findIndicatorFromCacheWithTracing := main_usecases.NewFindIndicatorFromCacheWithTracing(&indicatorCachedDatabaseGateway)
	saveIndicatorFromCacheWithTracing := main_usecases.NewSaveIndicatorFromCacheWithTracing(&indicatorCachedDatabaseGateway)

	cachedFindIndicatorById := main_usecases.NewFindIndicatorByIdWithCache(
		findIndicatorById,
		findIndicatorFromCacheWithTracing,
		saveIndicatorFromCacheWithTracing,
	)

	return main_usecases.NewFindIndicatorByIdWithCacheAndLocked(cachedFindIndicatorById)
}
