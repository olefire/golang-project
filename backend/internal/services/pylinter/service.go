package pylinter

import (
	"backend/internal/models/linter"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

var (
	Python linter.ProgrammingLanguage = "python"
)

type PylintLinter struct{}

type Service struct {
	linter *PylintLinter
}

func NewClient(linter *PylintLinter) *Service {
	return &Service{
		linter: linter,
	}
}

func (s *Service) LintCode(sourceFile linter.SourceFile) ([]linter.LintCodeIssue, error) {
	content, stderr, err := s.linter.lintFile(sourceFile)

	if stderr != "" {
		fmt.Printf("lintFile stderr: %s\n", stderr)
	}

	if err != nil {
		return nil, err
	}

	var issues []linter.LintCodeIssue
	err = json.Unmarshal([]byte(content), &issues)
	if err != nil {
		return nil, err
	}

	return issues, err
}

func (pylint *PylintLinter) lintFile(file linter.SourceFile) (stdout string, stderr string, err error) {
	cmd := exec.Command("pylint",
		"-f", "json",
		"--from-stdin", "file.py",
	)

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return "", "", err
	}

	_, err = pipe.Write([]byte(file.Code))
	if err != nil {
		return "", "", err
	}

	if err = pipe.Close(); err != nil {
		return "", "", err
	}

	// Substitute process stderr/stdout buffers
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err = cmd.Run()
	var e *exec.ExitError
	if !errors.As(err, &e) {
		fmt.Printf("unexpected error code: %s", err)
		return "", "", err
	}

	return outBuffer.String(), errBuffer.String(), nil
}
