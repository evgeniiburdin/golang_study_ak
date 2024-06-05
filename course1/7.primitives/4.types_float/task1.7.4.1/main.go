package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypotenuse(3, 4))
}

func hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
