package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoGo/internal/models"
)

type UserManagement interface {
	CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}

type PasteManagement interface {
	CreatePaste(ctx context.Context, paste *models.Paste) (*mongo.InsertOneResult, error)
	GetBatch(ctx context.Context) ([]models.Paste, error)
}
