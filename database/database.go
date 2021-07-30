package database

/**
 * This database mainly set up mongo db
 *
 * @author: Boris
 * @version: 2021-07-08
 *
 */

import (
	"context"
	"url-shortener/config"
	"url-shortener/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UrlCollection *mongo.Collection

//Setup mongo connection
func Setup() {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.Env.DBURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Error.Fatalf("Setup MongoDB connect error %+v\n", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Error.Fatalf("Setup MongoDB ping error %+v\n", err)
	}

	logger.Info.Printf("Connected to MongoDB")

	// Set database and collection
	UrlCollection = client.Database(config.Env.DBName).Collection(config.Env.UrlInfoCollectionName)
	if UrlCollection == nil {
		logger.Error.Fatalf("collection: %s is nil", config.Env.UrlInfoCollectionName)
	}
}
