package paste

import (
	"context"
	"fmt"
	"golang-project/internal/models"
)

type Repository interface {
	CreatePaste(ctx context.Context, paste *models.Paste) (string, error)
	GetBatch(ctx context.Context) ([]models.Paste, error)
	GetPasteById(ctx context.Context, id string) (*models.Paste, error)
	DeletePaste(ctx context.Context, id string) error
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

func (s *Service) GetPasteById(ctx context.Context, id string) (*models.Paste, error) {
	paste, err := s.PasteRepo.GetPasteById(ctx, id)

	if err != nil {
		return nil, err
	}

	return paste, err
}

func (s *Service) DeletePaste(ctx context.Context, id string) error {
	err := s.PasteRepo.DeletePaste(ctx, id)

	return err
}
