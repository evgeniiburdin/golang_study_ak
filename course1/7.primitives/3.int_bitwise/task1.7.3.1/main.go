package main

import "fmt"

func main() {
	var a int = 5
	var b int = 1

	fmt.Printf("%d & %d = %d\n", a, b, bitwiseAnd(a, b))
	fmt.Printf("%d | %d = %d\n", a, b, bitwiseOr(a, b))
	fmt.Printf("%d ^ %d = %d\n", a, b, bitwiseXor(a, b))
	fmt.Printf("%d << %d = %d\n", a, b, bitwiseLeftShift(a, b))
	fmt.Printf("%d >> %d = %d\n", a, b, bitwiseRightShift(a, b))
}

func bitwiseAnd(a, b int) int {
	return a & b
}

func bitwiseOr(a, b int) int {
	return a | b
}

func bitwiseXor(a, b int) int {
	return a ^ b
}

func bitwiseLeftShift(a, b int) int {
	return a << b
}

func bitwiseRightShift(a, b int) int {
	return a >> b
}
