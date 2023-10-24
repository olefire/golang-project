package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoGo/internal/models"
)

type UserManagement interface {
	CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	GetUsersInfo(ctx context.Context) ([]models.User, error)
}
