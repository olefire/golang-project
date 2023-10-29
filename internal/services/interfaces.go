package services

import (
	"context"
	"mongoGo/internal/models"
)

type UserManagement interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}

type PasteManagement interface {
	CreatePaste(ctx context.Context, paste *models.Paste) (string, error)
	GetBatch(ctx context.Context) ([]models.Paste, error)
}
