package main

import (
	"bufio"
	"bytes"
)

func getScanner(b *bytes.Buffer) *bufio.Scanner {
	scanner := bufio.NewScanner(b)
	return scanner
}
