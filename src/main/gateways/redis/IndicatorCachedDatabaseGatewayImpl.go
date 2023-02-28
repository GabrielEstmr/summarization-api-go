package main_gateways_redis

import (
	main_domains "mpindicator/main/domains"
	main_gateways_redis_documents "mpindicator/main/gateways/redis/documents"
	main_gateways_redis_repositories "mpindicator/main/gateways/redis/repositories"
)

type IndicatorCachedDatabaseGatewayImpl struct {
	cachedIndicatorRepository *main_gateways_redis_repositories.CachedIndicatorRepository
}

func NewIndicatorCachedDatabaseGatewayImpl(
	cachedIndicatorRepository *main_gateways_redis_repositories.CachedIndicatorRepository,
) *IndicatorCachedDatabaseGatewayImpl {
	return &IndicatorCachedDatabaseGatewayImpl{cachedIndicatorRepository}
}

func (thisGateway *IndicatorCachedDatabaseGatewayImpl) Save(
	indicator main_domains.Indicator) (main_domains.Indicator, error) {

	indicatorDocument, err := thisGateway.cachedIndicatorRepository.Save(
		main_gateways_redis_documents.NewCachedIndicatorDocument(indicator))
	return indicatorDocument.ToDomain(), err
}

func (thisGateway *IndicatorCachedDatabaseGatewayImpl) FindById(id string) (main_domains.Indicator, error) {
	indicatorDocument, err := thisGateway.cachedIndicatorRepository.FindById(id)
	return indicatorDocument.ToDomain(), err
}
