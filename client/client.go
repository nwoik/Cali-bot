package client

import (
	"calibot/globals"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() (*mongo.Client, error) {
	pswd := globals.MONGO_PASS
	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://mongo:"+pswd+"@viaduct.proxy.rlwy.net:58839/?tlsCertificateKeyFilePassword="+pswd))

	if err != nil {
		log.Println("error connecting to db", err)
		return nil, err
	}

	log.Println("successfully connected")

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Println("ping failed")
		return nil, err
	}

	return mongoClient, nil
}
