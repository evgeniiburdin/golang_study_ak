package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var i int = 6
	fmt.Println(Factorial(&i))

	var str string = "Я не мил, и не женили меня."
	fmt.Println(isPalindrome(&str))

	var nums []int = []int{9, 1, 6, 7, 1, 2, 9, 9}
	var target int = 9
	fmt.Println(CountOccurrences(&nums, &target))

	str2 := "eciuj egnarO"
	fmt.Println(ReverseString(&str2))
}

func Factorial(n *int) int {
	fac := 1
	for i := 1; i < *n+1; i++ {
		fac *= i
	}
	return fac
}

func isPalindrome(str *string) bool {
	tempStr := strings.ToLower(*str)
	words := strings.Split(tempStr, " ")
	runes := make([]rune, 0, len(*str))
	for _, word := range words {
		for _, runee := range word {
			runes = append(runes, runee)
		}
	}
	for r := 0; r < len(runes); r++ {
		if !unicode.In(runes[r], unicode.Cyrillic) {
			runes = append(runes[:r], runes[r+1:]...)
			r--
		}
	}
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func CountOccurrences(numbers *[]int, target *int) int {
	oc := 0
	for i := 0; i < len(*numbers); i++ {
		if (*numbers)[i] == *target {
			oc++
		}
	}
	return oc
}

func ReverseString(str *string) string {
	runes := []rune(*str)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
