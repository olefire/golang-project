package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-project/internal/config"
	controllerhttp "golang-project/internal/controller/http"
	PasteRepo "golang-project/internal/repository/paste"
	UserRepo "golang-project/internal/repository/user"
	PasteService "golang-project/internal/services/paste"
	UserService "golang-project/internal/services/user"
	"golang-project/pkg/middleware"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database(cfg.Database)

	userCollection := db.Collection(cfg.UserCollection)
	pasteCollection := db.Collection(cfg.PasteCollection)

	userRepo := UserRepo.NewUserRepository(userCollection)
	pasteRepo := PasteRepo.NewPasteRepository(pasteCollection)

	userService := UserService.NewService(UserService.Deps{UserRepo: userRepo})
	pasteService := PasteService.NewService(PasteService.Deps{PasteRepo: pasteRepo})

	ctr := controllerhttp.NewController(controllerhttp.UserService{UserManagement: userService},
		controllerhttp.PasteService{PasteManagement: pasteService})

	router := ctr.NewRouter()

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)

	err = http.ListenAndServe(cfg.Port, middleware.PanicRecovery(middleware.LogRequest(handler)))
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
