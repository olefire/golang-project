package services

import (
	"lint-service/internal/models"
)

type LinterManagement interface {
	LintCode(file models.SourceFile) ([]models.LintResult, error)
}
