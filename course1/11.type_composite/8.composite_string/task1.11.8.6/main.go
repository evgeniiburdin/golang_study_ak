package main

import "strings"

func CountVowels(str string) int {
	runes := []rune(str)

	vowels := "aeiou"
	count := 0

	for _, r := range runes {
		if strings.ContainsRune(vowels, r) {
			count++
		}
	}
	return count
}
