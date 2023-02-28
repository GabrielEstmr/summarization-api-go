package main_gateways_mongodb_repositories

import (
	"context"
	"fmt"

	main_configurations_mongodb "mpindicator/main/configurations/mongodb"
	main_gateway_mongodb_documents "mpindicator/main/gateways/mongodb/documents"

	"go.mongodb.org/mongo-driver/mongo"
)

const USERS_COLLECTION_NAME = "users"

type UserRepository struct {
	database *mongo.Database
}

func NewUserRepository() *UserRepository {
	return &UserRepository{database: main_configurations_mongodb.GetDatabaseBean()}
}

func (userRepository *UserRepository) Save(ctx context.Context,
	user main_gateway_mongodb_documents.UserDocument) (string, error) {

	collection := userRepository.database.Collection(USERS_COLLECTION_NAME)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(result.InsertedID), nil

}
