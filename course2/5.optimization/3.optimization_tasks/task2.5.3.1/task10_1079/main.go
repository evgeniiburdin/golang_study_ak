package main

func numTilePossibilities(tiles string) int {
	count := make(map[rune]int)
	for _, tile := range tiles {
		count[tile]++
	}
	return dfs(count)
}

func dfs(count map[rune]int) int {
	sum := 0
	for tile, num := range count {
		if num > 0 {
			count[tile]--
			sum += 1 + dfs(count)
			count[tile]++
		}
	}
	return sum
}
