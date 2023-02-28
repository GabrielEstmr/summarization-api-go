package main_gateways_mongodb_repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	main_configurations_mongodb "mpindicator/main/configurations/mongodb"
	main_gateway_mongodb_documents "mpindicator/main/gateways/mongodb/documents"

	"go.mongodb.org/mongo-driver/mongo"
)

const INDICATORS_COLLECTION_NAME = "indicators"

type IndicatorRepository struct {
	database *mongo.Database
}

func NewIndicatorRepository() *IndicatorRepository {
	return &IndicatorRepository{database: main_configurations_mongodb.GetDatabaseBean()}
}

const IDX_INDICATOR_MONGO_ID = "_id"

func (thisRepository *IndicatorRepository) FindById(ctx context.Context,
	id string) (main_gateway_mongodb_documents.IndicatorDocument, error) {

	collection := thisRepository.database.Collection(INDICATORS_COLLECTION_NAME)
	var result main_gateway_mongodb_documents.IndicatorDocument
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.D{{IDX_INDICATOR_MONGO_ID, objectId}}
	err2 := collection.FindOne(ctx, filter).Decode(&result)
	if err2 != nil && err2 != mongo.ErrNoDocuments {
		return result, err2
	}
	return result, nil
}

func (thisRepository *IndicatorRepository) Save(ctx context.Context,
	indicatorDocument main_gateway_mongodb_documents.IndicatorDocument) (string, error) {

	collection := thisRepository.database.Collection(INDICATORS_COLLECTION_NAME)
	result, err := collection.InsertOne(ctx, indicatorDocument)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(result.InsertedID), nil
}

func (thisRepository *IndicatorRepository) SaveAll(ctx context.Context,
	indicatorDocuments []main_gateway_mongodb_documents.IndicatorDocument) ([]string, error) {

	collection := thisRepository.database.Collection(INDICATORS_COLLECTION_NAME)
	var newIndicators []interface{}
	for _, value := range indicatorDocuments {
		newIndicators = append(newIndicators, value)
	}

	result, err := collection.InsertMany(ctx, newIndicators)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, value := range result.InsertedIDs {
		ids = append(ids, fmt.Sprint(value))
	}

	return ids, nil
}
