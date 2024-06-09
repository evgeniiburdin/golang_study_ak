package main

import "strings"

func filterSentence(sentence string, filter map[string]bool) string {
	words := strings.Fields(sentence)
	var resultWords []string

	for _, word := range words {
		if !filter[word] {
			resultWords = append(resultWords, word)
		}
	}
	return strings.Join(resultWords, " ")
}
