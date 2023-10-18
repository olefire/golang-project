package service

import (
	"context"
	"mongoGo/internal/models"
)

type Repository interface {
	Create(c context.Context, user *models.User) error
	Fetch(c context.Context) ([]models.User, error)
	GetByEmail(c context.Context, email string) (models.User, error)
	GetByID(c context.Context, id string) (models.User, error)
}
