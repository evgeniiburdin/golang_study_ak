package main

import (
	"fmt"

	"sort"
)

func main() {
	fmt.Println(FindMinAndMax(7, 9, 1, 6, 3, 0, 2, 4, 5))
}

func FindMinAndMax(n ...int) (int, int) {
	sort.Ints(n)
	return n[0], n[len(n)-1]
}
