package main

import "fmt"

func main() {
	fmt.Println(*Add(-5, 8))
	fmt.Println(*Max([]int{1, 9, 10, 15, 24, 3, 98}))
	fmt.Println(*IsPrime(37))
	fmt.Println(*ConcatenateStrings([]string{"hello ", "world", "!"}))
}

func Add(a, b int) *int {
	sum := a + b
	return &sum
}

func Max(numbers []int) *int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return &max
}

func IsPrime(number int) *bool {
	var isPrime bool = true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}
	return &isPrime
}

func ConcatenateStrings(strs []string) *string {
	var conc string = ""
	for _, str := range strs {
		conc += str
	}
	return &conc
}
