package linter

import (
	"fmt"
	"lint-service/internal/linters"
	"lint-service/internal/models"
	"sync"
)

type Service struct {
	linter []linters.Linter
}

func NewClient(linter []linters.Linter) *Service {
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
			var err error
			lintResult[i], err = linter.LintFile(sourceFile)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()

	return lintResult, nil
}
