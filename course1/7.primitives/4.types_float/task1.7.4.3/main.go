package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 12.34567
	var b float64 = 12.38
	fmt.Printf("a = %f\n", a)
	fmt.Printf("b = %f\n", b)
	fmt.Println(CompareRoundedValues(a, b, 4))
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	epsilon := math.Pow(0.1, float64(decimalPlaces))
	roundedA := math.Round(a*math.Pow(10, float64(decimalPlaces))) / math.Pow(10, float64(decimalPlaces))
	roundedB := math.Round(b*math.Pow(10, float64(decimalPlaces))) / math.Pow(10, float64(decimalPlaces))
	isEqual = math.Abs(roundedA-roundedB) < epsilon
	difference = math.Abs(roundedA - roundedB)
	return isEqual, difference
}
