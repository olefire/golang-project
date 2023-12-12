package linter

import (
	"lint-service/internal/models"
	"lint-service/internal/services"
)

var (
	Python models.ProgrammingLanguage = "python"
)

type Service struct {
	linter []services.Linter
}

func NewClient(linter []services.Linter) *Service {
	return &Service{
		linter: linter,
	}
}

func (s *Service) LintCode(sourceFile models.SourceFile) ([]models.LintResult, error) {
	var lintResult []models.LintResult
	for _, linter := range s.linter {
		issues, err := linter.LintFile(sourceFile)

		if err != nil {
			return []models.LintResult{}, err
		}

		lintResult = append(lintResult, issues)
	}
	return lintResult, nil
}
