package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sin(90))
	fmt.Println(Cos(90))
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}
