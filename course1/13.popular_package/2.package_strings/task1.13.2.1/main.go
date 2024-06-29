package main

import "strings"

func CountWordsInText(txt string, words []string) map[string]int {
	textWords := strings.Fields(txt)
	textWordsCount := map[string]int{}
	for _, word := range textWords {
		for _, w := range words {
			if strings.Contains(strings.ToLower(word), strings.ToLower(w)) {
				textWordsCount[strings.ToLower(w)]++
			}
		}
	}
	return textWordsCount
}
