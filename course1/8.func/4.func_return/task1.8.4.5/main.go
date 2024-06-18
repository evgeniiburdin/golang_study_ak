package main

import (
	"fmt"
	"math"
)

func main() {
	initValue, finalValue := 10.0, 15.0
	fmt.Printf("init value: %f, final value: %f, percent change: %.2f%%\n",
		initValue, finalValue, CalculatePercentageChange(initValue, finalValue))
}

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	change := (math.Abs(finalValue-initialValue) / initialValue) * 100
	return math.Round(change*100) / 100
}
