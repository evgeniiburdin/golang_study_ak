package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Floor(95.1351357))
}

func Floor(x float64) float64 {
	return math.Floor(x)
}
