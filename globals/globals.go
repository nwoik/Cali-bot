package globals

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	SERVER_HOST string
	TOKEN       string
	CLIENT      *mongo.Client
)

func InitConfig() {
	TOKEN = os.Getenv("CALIBOT_TOKEN")
	SERVER_HOST = os.Getenv("MONGO_URL")
	var err error
	CLIENT, err = NewMongoClient()
	if err != nil {
		log.Fatalln("Couldn't connect to the database", err)
	}
}

func NewMongoClient() (*mongo.Client, error) {
	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(SERVER_HOST))

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
