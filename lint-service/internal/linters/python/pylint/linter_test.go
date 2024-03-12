package pylint

import (
	"lint-service/internal/models"
	"os"
	"path"
	"testing"
)

var pyLint = Linter{}

func TestService_LintCode(t *testing.T) {
	projectRoot := "../../../.."
	fpath := path.Join(projectRoot, "code/test_file.py")
	content, err := os.ReadFile(fpath)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	file := models.SourceFile{
		Code:     string(content),
		Language: models.Python,
	}

	issues, err := pyLint.LintFile(file)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	t.Log(issues)
}
