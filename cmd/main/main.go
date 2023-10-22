package main

import (
	"context"
	"github.com/rs/cors"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"log"
	controllerhttp "mongoGo/internal/controller/http"
	"mongoGo/internal/repository"
	"mongoGo/internal/services"
	"mongoGo/pkg/client/mongo"
	"mongoGo/pkg/handlers"
	"mongoGo/pkg/middleware"
	"net/http"
)

func main() {
	client, err := mongo.NewMongoDatabase(context.Background())
	if err != nil {
		log.Fatalf("Failed connect to database: %v", err)
	}
	defer func(ctx context.Context, client *mongo2.Client) {
		err := mongo.CloseMongoDBConnection(ctx, client)
		if err != nil {
			log.Fatalf("Database disconnect error")
		}
	}(context.Background(), client)

	db := client.Database(handlers.DotEnvVariable("DATABASE"))

	repo := repository.NewUserRepository(db, handlers.DotEnvVariable("COLLECTION"))

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
