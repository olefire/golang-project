package pylint

import (
	"encoding/json"
	"fmt"
	"lint-service/internal/models"
	"lint-service/pkg"
)

type Linter struct{}

var (
	Pylint models.Linter = "Pylint"
)

func (l *Linter) LintFile(file models.SourceFile) (models.LintResult, error) {

	content, stderr, err := pkg.Execute(file.Code, "pylint", "-f", "json",
		"--from-stdin", "file.py")

	if stderr != "" {
		fmt.Printf("lintFile stderr: %s\n", stderr)
	}

	if err != nil {
		return models.LintResult{}, err
	}

	var issues []models.LintCodeIssue
	err = json.Unmarshal([]byte(content), &issues)
	if err != nil {
		return models.LintResult{}, err
	}

	return models.LintResult{Issues: issues, Linter: Pylint}, nil
}
