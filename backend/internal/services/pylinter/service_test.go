package pylinter

import (
	"backend/internal/models/linter"
	"os"
	"path"
	"testing"
)

func TestService_LintCode(t *testing.T) {
	service := NewClient(&PylintLinter{})

	projectRoot := "../../.."
	fpath := path.Join(projectRoot, "code/test_file.py")
	content, err := os.ReadFile(fpath)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	file := linter.SourceFile{
		Code:     string(content),
		Language: Python,
	}

	issues, err := service.LintCode(file)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	if len(issues) == 0 {
		t.Errorf("expected more than at least one issue")
	}

	t.Logf("found %d issues", len(issues))
}
