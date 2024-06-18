package main

import (
	"bufio"

	"fmt"

	"os"
)

func ReadString(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("couldn't open file %s: %#v\n", filePath, err)
		return ""
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("couldn't close file %s: %#v\n", filePath, err)
			return
		}
	}()
	reader := bufio.NewReader(file)
	content, err := reader.ReadString('/')
	if err != nil {
		fmt.Printf("couldn't read line from file %s: %#v\n", filePath, err)
		return ""
	}
	return content
}
