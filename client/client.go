package client

import (
	"calibot/globals"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	pswd := globals.MONGO_PASS
	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://mongo:"+pswd+"@viaduct.proxy.rlwy.net:58839/?tlsCertificateKeyFilePassword="+pswd))

	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	log.Println("successfully connected")

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed")
	}

	return mongoClient
}
