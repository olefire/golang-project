package main

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-project/internal/config"
	controllerhttp "golang-project/internal/controller/http"
	PasteRepo "golang-project/internal/repository/paste"
	UserRepo "golang-project/internal/repository/user"
	"golang-project/internal/services/linter"
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

	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer func(cli *client.Client) {
		err := cli.Close()
		if err != nil {
			panic(err)
		}
	}(dockerClient)

	dClient := linter.NewClient(dockerClient)

	clientOptions := options.Client().ApplyURI(cfg.MongoURL)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := mongoClient.Database(cfg.Database)

	userCollection := db.Collection(cfg.UserCollection)
	pasteCollection := db.Collection(cfg.PasteCollection)

	userRepo := UserRepo.NewUserRepository(userCollection)
	pasteRepo := PasteRepo.NewPasteRepository(pasteCollection)

	userService := UserService.NewService(UserService.Deps{UserRepo: userRepo})
	pasteService := PasteService.NewService(PasteService.Deps{PasteRepo: pasteRepo})

	ctr := controllerhttp.NewController(controllerhttp.UserService{UserManagement: userService},
		controllerhttp.PasteService{PasteManagement: pasteService},
		controllerhttp.LinterService{Linter: dClient})

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
