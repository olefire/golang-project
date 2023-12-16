package linters

import "lint-service/internal/models"

type Linter interface {
	LintFile(file models.SourceFile) (models.LintResult, error)
}
