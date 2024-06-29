package main

func countUniqueUTF8Chars(s string) int {
	runes := []rune(s)
	uniqueUTF8 := make(map[rune]struct{})
	for _, r := range runes {
		if _, ok := uniqueUTF8[r]; !ok {
			uniqueUTF8[r] = struct{}{}
		}
	}
	return len(uniqueUTF8)
}
