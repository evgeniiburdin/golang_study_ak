package main

import "fmt"

func main() {
	div, rem := DivideAndRemainder(10, 5)
	fmt.Printf("Частное: %d, Остаток: %d\n", div, rem)
}

func DivideAndRemainder(a, b int) (int, int) {
	if b == 0 {
		return 0, 0
	}
	return a / b, a % b
}
