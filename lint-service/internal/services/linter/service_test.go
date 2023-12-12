package linter

import (
	"lint-service/internal/models"
	"lint-service/internal/services"
	"lint-service/internal/services/metrics"
	"lint-service/internal/services/pylint"
	"os"
	"path"
	"testing"
)

func TestService_GetMetrics(t *testing.T) {
	pyLint := pylint.Linter{}
	pyMetrics := metrics.Linter{}

	linters := []services.Linter{&pyLint, &pyMetrics}

	linterService := NewClient(linters)

	projectRoot := "../../.."
	fpath := path.Join(projectRoot, "code/test_file.py")
	content, err := os.ReadFile(fpath)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	file := models.SourceFile{
		Code:     string(content),
		Language: Python,
	}

	issues, err := linterService.LintCode(file)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	t.Log("found issues", issues)
}
