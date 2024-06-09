package main

func RemoveExtraMemory(xs []int) []int {
	xs = append([]int(nil), xs...)
	return xs
}
