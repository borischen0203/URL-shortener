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
	"log"
	"url-shortener/config"
	"url-shortener/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var (
	collectionurlInfo string = "url_info"
	databaseUrl       string
	UrlCollection     *mongo.Collection
)

// var UrlCollection *mongo.Collection

//Setup mongo connection
func Setup() {
	cs, err := connstring.ParseAndValidate(config.Env.MongoURI)
	if err != nil {
		logger.Error.Fatalf("SetupMongoDB parse MongoURI error %+v\n", err)
	}
	databaseUrl = cs.Database

	// Set client options
	clientOptions := options.Client().ApplyURI(config.Env.MongoURI)
	// clientOptions := options.Client().ApplyURI("mongodb+srv://root:root@cluster0.qfx1p.mongodb.net/short-url?retryWrites=true&w=majority")
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27016")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	// UrlCollection = client.Database(databaseUrl).Collection(collectionurlInfo)
	UrlCollection = client.Database("url_database").Collection(collectionurlInfo)
	if UrlCollection == nil {
		log.Fatal("collection is nil")
	}
	log.Println("SetupDB successful")
}
