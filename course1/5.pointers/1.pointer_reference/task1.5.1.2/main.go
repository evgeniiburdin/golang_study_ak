package main

import "fmt"

func main() {
	var i int = 3
	mutate(&i)

	fmt.Println(i)

	str := "eciuj egnarO"
	ReverseString(&str)

	fmt.Println(str)
}

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	runes := []rune(*str)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	*str = string(runes)
}
