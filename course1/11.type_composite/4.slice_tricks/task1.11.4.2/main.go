package main

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	xs = append(xs[:idx+1], append(x, xs[idx+1:]...)...)
	return xs
}
