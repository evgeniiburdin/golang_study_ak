package main

import "strings"

func filterSentence(sentence string, filter map[string]struct{}) string {
	words := strings.Fields(sentence)
	var resultWords []string

	for _, word := range words {
		if _, ok := filter[word]; !ok {
			resultWords = append(resultWords, word)
		}
	}

	return strings.Join(resultWords, " ")
}
