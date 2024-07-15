package main

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		if isRussianLetter(r) {
			counts[r]++
		}
	}
	return counts
}

func isRussianLetter(r rune) bool {
	return r >= 'А' && r <= 'Я' || r >= 'а' && r <= 'я'
}
