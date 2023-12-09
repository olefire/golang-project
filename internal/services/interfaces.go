package services

import (
	"context"
	"golang-project/internal/models"
	"io"
)

type UserManagement interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type PasteManagement interface {
	CreatePaste(ctx context.Context, paste *models.Paste) (string, error)
	GetBatch(ctx context.Context) ([]models.Paste, error)
	GetPasteById(ctx context.Context, id string) (*models.Paste, error)
	DeletePaste(ctx context.Context, id string) error
}

type Linter interface {
	LintCode(ctx context.Context) io.ReadCloser
}
