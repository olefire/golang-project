package linter

import (
	"lint-service/internal/linters/python/metrics"
	"lint-service/internal/linters/python/pylint"
	"lint-service/internal/models"
	"lint-service/internal/services"
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
		Language: models.Python,
	}

	issues, err := linterService.LintCode(file)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	t.Log("found issues", issues)
}
