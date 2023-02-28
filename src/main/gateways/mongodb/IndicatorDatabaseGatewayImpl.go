package main_gateways_mongodb

import (
	"context"
	main_domains "mpindicator/main/domains"
	main_gateway_mongodb_documents "mpindicator/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "mpindicator/main/gateways/mongodb/repositories"
)

type IndicatorDatabaseGatewayImpl struct {
	indicatorRepository *main_gateways_mongodb_repositories.IndicatorRepository
}

func NewIndicatorDatabaseGatewayImpl(indicatorRepository *main_gateways_mongodb_repositories.IndicatorRepository) *IndicatorDatabaseGatewayImpl {
	return &IndicatorDatabaseGatewayImpl{indicatorRepository}
}

func (thisGateway *IndicatorDatabaseGatewayImpl) FindById(ctx context.Context, id string) (main_domains.Indicator, error) {
	indicatorDocument, err := thisGateway.indicatorRepository.FindById(ctx, id)
	return indicatorDocument.ToDomain(), err
}

func (thisGateway *IndicatorDatabaseGatewayImpl) Save(ctx context.Context, indicator main_domains.Indicator) (string, error) {
	indicatorDocument := main_gateway_mongodb_documents.NewIndicatorDocument(indicator)
	id, err := thisGateway.indicatorRepository.Save(ctx, indicatorDocument)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (thisGateway *IndicatorDatabaseGatewayImpl) SaveAll(ctx context.Context, indicators []main_domains.Indicator) ([]string, error) {
	var indicatorDocuments []main_gateway_mongodb_documents.IndicatorDocument
	for _, value := range indicators {
		indicatorDocuments = append(
			indicatorDocuments, main_gateway_mongodb_documents.NewIndicatorDocument(value))
	}
	ids, err := thisGateway.indicatorRepository.SaveAll(ctx, indicatorDocuments)
	if err != nil {
		return make([]string, 0), err
	}
	return ids, nil
}
