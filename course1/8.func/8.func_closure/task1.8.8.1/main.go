package main

import "fmt"

func main() {
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func createCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
