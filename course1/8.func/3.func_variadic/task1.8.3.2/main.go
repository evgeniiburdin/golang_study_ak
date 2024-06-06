package main

import (
	"fmt"
	"strings"
)

func main() {
	result := ConcatenateStrings("-", "hello", "world", "how", "are", "you")
	fmt.Println(result)
}

func ConcatenateStrings(sep string, str ...string) string {
	var evenStrings []string
	var oddStrings []string

	for i, el := range str {
		if i%2 == 0 {
			evenStrings = append(evenStrings, el)
		} else {
			oddStrings = append(oddStrings, el)
		}
	}

	evenResult := strings.Join(evenStrings, sep)
	oddResult := strings.Join(oddStrings, sep)

	evenResult = strings.TrimSuffix(evenResult, sep)
	oddResult = strings.TrimSuffix(oddResult, sep)

	return fmt.Sprintf("even: %s, odd: %s", evenResult, oddResult)
}
