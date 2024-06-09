package task1_11_7_1

import "strings"

func countWordOccurrences(text string) map[string]int {
	words := strings.Fields(text)
	occurrences := map[string]int{}
	for _, word := range words {
		occurrences[strings.ToLower(word)]++
	}
	return occurrences
}
