package service

import (
	"errors"
	"go-testDB/internal/models"
)

var users []models.User

func CreateUser(user models.User) (*models.User, error) {
	for _, existingUser := range users {
		if existingUser.Email == user.Email {
			return nil, errors.New("email already exists")
		}
	}

	users = append(users, user)

	return &user, nil
}
