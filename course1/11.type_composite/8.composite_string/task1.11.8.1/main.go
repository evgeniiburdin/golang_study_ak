package main

import "unicode/utf8"

func countBytes(s string) int {
	return len(s)
}

func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}
