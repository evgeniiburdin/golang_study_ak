package main

import "strings"

func createUniqueText(text string) string {
	words := strings.Fields(text)
	occurrences := make(map[string]struct{})
	var uniqueWords []string
	for _, word := range words {
		if _, ok := occurrences[word]; !ok {
			occurrences[word] = struct{}{}
			uniqueWords = append(uniqueWords, word)
		}
	}
	return strings.Join(uniqueWords, " ")
}
