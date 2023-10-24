package main

import (
	"context"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	controllerhttp "mongoGo/internal/controller/http"
	"mongoGo/internal/repository"
	"mongoGo/internal/services"
	"mongoGo/pkg/handlers"
	"mongoGo/pkg/middleware"
	"net/http"
)

func main() {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(handlers.DotEnvVariable("MONGO_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if client.Ping(ctx, nil) != nil {
		log.Fatal(err)
	}
	defer func() {
		if client.Disconnect(ctx) != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database(handlers.DotEnvVariable("DATABASE"))

	collection := db.Collection(handlers.DotEnvVariable("COLLECTION"))

	repo := repository.NewUserRepository(collection)

	service := services.NewService(services.Deps{
		Repo: repo,
	})

	ctr := controllerhttp.NewController(controllerhttp.Services{
		UserManagement: service,
	})

	router := ctr.NewRouter()

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)

	err = http.ListenAndServe(":8080", middleware.LogRequest(handler))
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}

}
