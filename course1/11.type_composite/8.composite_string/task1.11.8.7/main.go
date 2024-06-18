package main

func ReplaceSymbols(s string, old rune, new rune) string {
	runes := []rune(s)
	for idx, r := range runes {
		if r == old {
			runes[idx] = new
		}
	}
	return string(runes)
}
