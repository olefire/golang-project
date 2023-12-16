package linter

import (
	"fmt"
	"lint-service/internal/models"
	"lint-service/internal/services"
	"sync"
)

type Service struct {
	linter []services.Linter
}

func NewClient(linter []services.Linter) *Service {
	return &Service{
		linter: linter,
	}
}

// LintCode todo handle errors
func (s *Service) LintCode(sourceFile models.SourceFile) ([]models.LintResult, error) {
	lintNumbers := len(s.linter)
	lintResult := make([]models.LintResult, lintNumbers)

	var wg sync.WaitGroup
	wg.Add(lintNumbers)

	for i, linter := range s.linter {
		i := i
		linter := linter
		go func() {
			defer wg.Done()
			currentLint, err := linter.LintFile(sourceFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			lintResult[i] = currentLint
		}()
	}

	wg.Wait()

	return lintResult, nil
}
