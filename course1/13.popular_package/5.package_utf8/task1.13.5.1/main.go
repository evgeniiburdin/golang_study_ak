package main

func countUniqueUTF8Chars(s string) int {
	runes := []rune(s)
	uniqueUTF8 := make(map[rune]bool)
	for _, r := range runes {
		if !uniqueUTF8[r] {
			uniqueUTF8[r] = false
		}
		uniqueUTF8[r] = true
	}
	counter := 0
	for _, v := range uniqueUTF8 {
		if v {
			counter++
		}
	}
	return counter
}
