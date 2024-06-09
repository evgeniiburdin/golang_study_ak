package main

import "strings"

func createUniqueText(text string) string {
	words := strings.Fields(text)
	occurrences := make(map[string]bool)
	var uniqueWords []string
	for _, word := range words {
		if !occurrences[word] {
			occurrences[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}
	return strings.Join(uniqueWords, " ")
}
