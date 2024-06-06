package main

import "fmt"

func main() {
	div, rem := DivideAndRemainder(10, 5)
	fmt.Printf("Частное: %d, Остаток: %d\n", div, rem)
}

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("check zero argument")
		}
	}()
	return a / b, a % b
}
