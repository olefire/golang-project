package main

import (
	"context"
	"fmt"
	"go-testDB/internal/repository"
)

func main() {
	client := repository.NewMongoDatabase()
	defer repository.CloseMongoDBConnection(&client)

	db := client.Database("test")

	userRepository := repository.NewUserRepository(*db, "users")

	//err := userRepository.Create(context.TODO(), &user)
	//if err != nil {
	//	return
	//}

	//users, err := userRepository.Fetch(context.TODO())
	//if err != nil {
	//	return
	//}
	//
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	user, err := userRepository.GetByEmail(context.TODO(), "test@mail.ru")
	if err != nil {
		return
	}

	fmt.Println(user)
}
