package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func Execute(code string, name string, arg ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command(name,
		arg...,
	)

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return "", "", err
	}

	go func() {
		defer pipe.Close()
		io.WriteString(pipe, code)
	}()

	// Substitute process stderr/stdout buffers
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err = cmd.Run()
	var e *exec.ExitError
	if err != nil && !errors.As(err, &e) {
		fmt.Printf("unexpected error code: %s", err)
		return "", "", err
	}
	return outBuffer.String(), errBuffer.String(), nil
}
