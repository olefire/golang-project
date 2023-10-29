package paste

import (
	"context"
	"fmt"
	"mongoGo/internal/models"
)

type Repository interface {
	CreatePaste(ctx context.Context, paste *models.Paste) (string, error)
	GetBatch(ctx context.Context) ([]models.Paste, error)
	GetPasteByTitle(ctx context.Context, title string) (*models.Paste, error)
}

type Deps struct {
	PasteRepo Repository
}

type Service struct {
	Deps
}

func NewService(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}

func (s *Service) CreatePaste(ctx context.Context, paste *models.Paste) (string, error) {
	insertedId, err := s.PasteRepo.CreatePaste(ctx, paste)
	if err != nil {
		return "", fmt.Errorf("can`t create paste")
	}
	return insertedId, err
}

func (s *Service) GetBatch(ctx context.Context) ([]models.Paste, error) {
	batch, err := s.PasteRepo.GetBatch(ctx)
	return batch, err
}
