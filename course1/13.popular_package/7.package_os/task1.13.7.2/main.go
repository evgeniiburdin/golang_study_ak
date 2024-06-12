package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func WriteFile(data io.Reader, fd io.Writer) error {
	if file, ok := fd.(*os.File); ok {
		dir := filepath.Dir(file.Name())
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("MkdirAll: %s", err)
		}
	} else {
		return fmt.Errorf("not a file desc: %#v", fd)
	}

	_, err := io.Copy(fd, data)
	if err != nil {
		return fmt.Errorf("io.Copy: %s", err)
	}
	return nil
}
