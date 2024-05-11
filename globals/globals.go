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
	SERVER_PORT string
	TOKEN       string
	MONGO_PASS  string
	CLIENT      *mongo.Client
)

func InitConfig() {
	TOKEN = os.Getenv("SQUIRE_TOKEN")
	MONGO_PASS = os.Getenv("MONGO_PASS")
	SERVER_PORT = "58839"
	SERVER_HOST = "mongodb://mongo:" + MONGO_PASS + "@viaduct.proxy.rlwy.net:" + SERVER_PORT + "/?tlsCertificateKeyFilePassword=" + MONGO_PASS
	var err error
	CLIENT, err = NewMongoClient()
	if err != nil {
		log.Fatalln("Couldn't connect to the database", err)
	}
}

func NewMongoClient() (*mongo.Client, error) {
	pswd := MONGO_PASS
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
