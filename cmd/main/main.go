package main

import (
	"context"
	"go-testDB/internal/models"
	"go-testDB/internal/repository"
)

func main() {
	client := repository.NewMongoDatabase()
	defer repository.CloseMongoDBConnection(&client)

	db := client.Database("test")

	userRepository := repository.NewUserRepository(*db, "users")

	user := models.User{ID: 1, FirstName: "Denis", LastName: "Hohlov", Email: "den4ik@mail.ru"}

	err := userRepository.Create(context.TODO(), &user)
	if err != nil {
		return
	}
}
