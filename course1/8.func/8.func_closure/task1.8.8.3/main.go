package main

import "fmt"

func main() {
	addTwo := adder(2)
	result := addTwo(3)
	fmt.Println(result)
}

func adder(initial int) func(value int) int {
	return func(value int) int {
		return initial + value
	}
}
