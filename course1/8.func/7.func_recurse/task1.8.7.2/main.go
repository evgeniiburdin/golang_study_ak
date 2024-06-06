package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(0))
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(6))
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
