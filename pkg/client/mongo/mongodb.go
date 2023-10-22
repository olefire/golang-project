package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongoGo/pkg/handlers"
)

func NewMongoDatabase(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(handlers.DotEnvVariable("MONGO_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, err
}

func CloseMongoDBConnection(ctx context.Context, client *mongo.Client) error {
	if client == nil {
		return errors.New("client is nil")
	}

	err := client.Disconnect(ctx)
	if err == nil {
		log.Println("Connection to MongoDB closed.")
	}

	return err
}
