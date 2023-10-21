package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongoGo/pkg/handlers"
)

func NewMongoDatabase() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(handlers.DotEnvVariable("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, err
}

func CloseMongoDBConnection(ctx context.Context, client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
