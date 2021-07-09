package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UrlCollection *mongo.Collection

//InitMongoDB function start to connect to mongo Database
func Setup() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://root:root@cluster0.qfx1p.mongodb.net/short-url?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	UrlCollection = client.Database("url_database").Collection("url_info")
	if UrlCollection == nil {
		log.Fatal("collection is nil")
	}
}
