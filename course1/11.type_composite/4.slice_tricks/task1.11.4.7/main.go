package main

import "fmt"

func main() {
	firstElement, slice := Pop([]int{1, 2, 3, 4, 5})
	fmt.Printf("Значение: %d, Новый срез: %v", firstElement, slice)
}

func Pop(xs []int) (int, []int) {
	return xs[0], xs[1:]
}
