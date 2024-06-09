package main

func FilterDividers(xs []int, divider int) []int {
	newSlice := make([]int, 0, len(xs))
	for _, el := range xs {
		if el%2 == 0 {
			newSlice = append(newSlice, el)
		}
	}
	return newSlice
}
