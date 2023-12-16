package metrics

import (
	"lint-service/internal/models"
	"os"
	"path"
	"testing"
)

var metric = Linter{}

func TestService_GetMetrics(t *testing.T) {

	projectRoot := "../../../.."
	fpath := path.Join(projectRoot, "code/test_file.py")
	content, err := os.ReadFile(fpath)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	file := models.SourceFile{
		Code:     string(content),
		Language: Python,
	}

	issues, err := metric.LintFile(file)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	t.Log("found issues", issues)
}
