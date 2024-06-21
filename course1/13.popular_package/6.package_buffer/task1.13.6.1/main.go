package main

import (
	"bufio"

	"bytes"
)

// getReader takes bytes.Buffer obj and returns bufio.Reader of the given obj
func getReader(b bytes.Buffer) bufio.Reader {
	reader := bufio.NewReader(&b)
	return *reader
}
