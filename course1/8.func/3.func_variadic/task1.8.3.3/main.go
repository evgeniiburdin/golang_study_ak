package main

import "fmt"

func main() {
	PrintNumbers(3, 4, 5, 6, 7, 8, 9, 0)
}

func PrintNumbers(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}
