package user

import (
	"backend/internal/models"
	"context"
	"fmt"
)

type Repository interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type Deps struct {
	UserRepo Repository
}

type Service struct {
	Deps
}

func NewService(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}

func (s *Service) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.UserRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("can`t get user: %w", err)
	}
	return user, err
}

func (s *Service) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.UserRepo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("can`t get users: %w", err)
	}
	return users, err
}

func (s *Service) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, err := s.UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can`t get user: %w", err)
	}
	return user, err
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	err := s.UserRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("can`t delete user: %w", err)
	}
	return err
}
