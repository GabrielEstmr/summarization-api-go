package main_configurations_mongodb

import (
	"context"
	"go.elastic.co/apm/module/apmmongo/v2"
	"log"
	main_configurations_yml "mpindicator/main/configurations/yml"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MSG_MONGO_BEAN_INITIALIZING = "Initializing mongo database bean."
const MSG_MONGO_BEAN_PINGED = "Successfully connected and pinged."
const MSG_MONGO_BEAN_CLOSING_CONNECTION = "Closing connection of mongo database bean."
const MSG_ERROR_TO_CONNECT_TO_DATABASE = "Application has been failed to connect to mongo database. URI: %s"
const MSG_ERROR_TO_PING_DATABASE = "Application has been failed to ping mongo database. URI: %s"

const MONGO_URI_NAME = "MongoDB.URI"
const MONGO_DATABASE_NAME = "MongoDB.DatabaseName"

var MongoDatabase *mongo.Database = nil
var once sync.Once

func GetDatabaseBean() *mongo.Database {
	once.Do(func() { // <-- atomic, does not allow repeating

		if MongoDatabase == nil {
			MongoDatabase = getDatabaseConnection()
		} // <-- thread safe

	})
	return MongoDatabase
}

func CloseConnection() {
	log.Println(MSG_MONGO_BEAN_CLOSING_CONNECTION)
	MongoDatabase.Client().Disconnect(context.TODO())
}

func getDatabaseConnection() *mongo.Database {

	log.Println(MSG_MONGO_BEAN_INITIALIZING)

	databaseUri := main_configurations_yml.GetBeanPropertyByName(MONGO_URI_NAME)
	mongoDatabaseName := main_configurations_yml.GetBeanPropertyByName(MONGO_DATABASE_NAME)

	// TODO: isolate this dependency from this config
	client, err := mongo.Connect(context.Background(),
		options.Client().SetMonitor(apmmongo.CommandMonitor()).ApplyURI(databaseUri))
	if err != nil {
		log.Fatalf(MSG_ERROR_TO_CONNECT_TO_DATABASE, databaseUri)
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf(MSG_ERROR_TO_PING_DATABASE, databaseUri)
		panic(err)
	}
	log.Println(MSG_MONGO_BEAN_PINGED)

	return client.Database(mongoDatabaseName)

}
