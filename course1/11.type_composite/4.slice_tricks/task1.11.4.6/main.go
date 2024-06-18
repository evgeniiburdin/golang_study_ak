package main

func InsertToStart(xs []int, x ...int) []int {
	return append(x, xs...)
}
