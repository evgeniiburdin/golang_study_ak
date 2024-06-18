package main

import "fmt"

func main() {
	var sum int
	var difference int
	var product int
	var quotient int
	var remainder int

	sum, difference, product, quotient, remainder = calculate(10, 3)
	fmt.Printf("sum = %d difference = %d product = %d quotient = %d remainder = %d", sum, difference,
		product, quotient, remainder)
}

func calculate(a int, b int) (int, int, int, int, int) {
	return a + b, a - b, a * b, a / b, a % b
}
