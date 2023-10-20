package main

import (
	"github.com/rs/cors"
	"log"
	"mongoGo/internal/repository"
	"mongoGo/internal/routes"
	"mongoGo/internal/services"
	"mongoGo/pkg/client/mongo"
	"mongoGo/pkg/middleware"
	"net/http"
)

func main() {

	client, err := mongo.NewMongoDatabase()
	if err != nil {
		log.Fatalf("Failed connect to database: %v", err)
	}

	db := client.Database("golang")

	repo := repository.NewUserRepository(db, "users")

	service := services.NewService(services.Deps{
		Repo: repo,
	})

	router := routes.Routes()

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
