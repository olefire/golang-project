package metrics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"lint-service/internal/models"
	"os/exec"
)

type Metrics struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Complexity int    `json:"complexity"`
	Rank       string `json:"rank"`
	Lineno     int    `json:"lineno"`
}

var (
	Python models.ProgrammingLanguage = "python"
	Radon  models.Linter              = "Radon"
)

type Linter struct{}

func (l *Linter) LintFile(file models.SourceFile) (models.LintResult, error) {
	cmd := exec.Command("radon",
		"cc", "-j", "-",
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

	var data map[string][]Metrics

	err = json.Unmarshal([]byte(outBuffer.String()), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return models.LintResult{}, err
	}
	issues := make([]models.LintCodeIssue, len(data["-"]))

	for i, metric := range data["-"] {
		issues[i].Message = fmt.Sprintf("Cyclomatic complexity of %s %s is : %d. Сode quality assessment: %s ", metric.Type, metric.Name, metric.Complexity, metric.Rank)
		issues[i].Line = metric.Lineno
	}

	return models.LintResult{Issues: issues, Linter: Radon}, nil
}
