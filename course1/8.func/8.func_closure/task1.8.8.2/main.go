package main

import "fmt"

func main() {
	m := multiplier(2.5)
	result := m(10)
	fmt.Println(result)
}

func multiplier(factor float64) func(float64) float64 {
	return func(x float64) float64 {
		x *= factor
		return x
	}
}
