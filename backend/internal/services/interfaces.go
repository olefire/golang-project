package services

import (
	"backend/internal/models"
	"backend/internal/models/linter"
	"context"
)

type UserManagement interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
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
	LintCode(sourceFile linter.SourceFile) ([]linter.LintCodeIssue, error)
}

type AuthManagement interface {
	SignUpUser(context.Context, *models.User) (string, error)
	SignInUser(context.Context, *models.SignInInput) (string, error)
	LogoutUser()
}
