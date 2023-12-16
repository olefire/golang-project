package pylint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"lint-service/internal/models"
	"os/exec"
)

type Linter struct{}

var (
	Pylint models.Linter = "Pylint"
)

func (l *Linter) LintFile(file models.SourceFile) (models.LintResult, error) {
	cmd := exec.Command("pylint",
		"-f", "json",
		"--from-stdin", "file.py",
	)

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return models.LintResult{}, err
	}

	_, err = pipe.Write([]byte(file.Code))
	if err != nil {
		return models.LintResult{}, err
	}

	if err = pipe.Close(); err != nil {
		return models.LintResult{}, err
	}

	// Substitute process stderr/stdout buffers
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err = cmd.Run()
	var e *exec.ExitError
	if err != nil && !errors.As(err, &e) {
		fmt.Printf("unexpected error code: %s", err)
		return models.LintResult{}, err
	}

	var issues []models.LintCodeIssue
	err = json.Unmarshal([]byte(outBuffer.String()), &issues)
	if err != nil {
		return models.LintResult{}, err
	}

	return models.LintResult{Issues: issues, Linter: Pylint}, nil
}
