package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecBin(binPath string, args ...string) string {
	cmd := exec.Command(binPath, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Sprintf("error executing %s: %s: %s", binPath, err.Error(), stderr.String())
	}

	return out.String()
}
