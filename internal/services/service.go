package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoGo/internal/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
}

type Deps struct {
	Repo Repository
}

type Service struct {
	Deps
}

func NewService(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}

func (s *Service) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	insertedId, err := s.Repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("can`t create user: %w", err)
	}
	return insertedId, err
}

func (s *Service) GetUsersInfo(ctx context.Context) ([]models.User, error) {
	users, err := s.Repo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("can`t get users: %w", err)
	}
	return users, err
}
