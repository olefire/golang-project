package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoGo/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
}

type PasteRepository interface {
	CreatePaste(ctx context.Context, paste *models.Paste) (*mongo.InsertOneResult, error)
	GetBatch(ctx context.Context) ([]models.Paste, error)
	GetPasteByTitle(ctx context.Context, title string) (*models.Paste, error)
}

type Deps struct {
	UserRepo  UserRepository
	PasteRepo PasteRepository
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
	insertedId, err := s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("can`t create user: %w", err)
	}
	return insertedId, err
}

func (s *Service) GetUsersInfo(ctx context.Context) ([]models.User, error) {
	users, err := s.UserRepo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("can`t get users: %w", err)
	}
	return users, err
}

func (s *Service) CreatePaste(ctx context.Context, paste *models.Paste) (*mongo.InsertOneResult, error) {
	insertedId, err := s.PasteRepo.CreatePaste(ctx, paste)
	if err != nil {
		return nil, fmt.Errorf("can`t create paste")
	}
	return insertedId, err
}

func (s *Service) GetBatchInfo(ctx context.Context) ([]models.Paste, error) {
	batch, err := s.PasteRepo.GetBatch(ctx)
	return batch, err
}
