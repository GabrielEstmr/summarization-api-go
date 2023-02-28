package main_gateways_mongodb

import (
	"context"
	main_domains "mpindicator/main/domains"
	main_gateway_mongodb_documents "mpindicator/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "mpindicator/main/gateways/mongodb/repositories"
)

type UserDatabaseGatewayImpl struct {
	userRepository *main_gateways_mongodb_repositories.UserRepository
}

func NewUserDatabaseGatewayImpl(userRepository *main_gateways_mongodb_repositories.UserRepository) *UserDatabaseGatewayImpl {
	return &UserDatabaseGatewayImpl{userRepository}
}

func (gateway *UserDatabaseGatewayImpl) Save(ctx context.Context, user main_domains.User) (string, error) {
	userDocument := main_gateway_mongodb_documents.NewUserDocument(user)
	id, err := gateway.userRepository.Save(ctx, userDocument)
	if err != nil {
		return "", err
	}
	return id, nil
}
