package main

import (
	"fmt"

	"io"

	"os"
	"path/filepath"
)

func WriteFile(data io.Reader, fd io.Writer) error {
	file, ok := fd.(*os.File)

	if !ok {
		return fmt.Errorf("not a file desc: %#v", fd)
	}

	dir := filepath.Dir(file.Name())
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("MkdirAll: %s", err)
	}

	_, err = io.Copy(fd, data)
	if err != nil {
		return fmt.Errorf("io.Copy: %s", err)
	}
	return nil
}
