package services

import (
	"context"
	"mongoGo/internal/models"
)

type UserManagement interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUsersInfo(ctx context.Context) ([]models.User, error)
}
