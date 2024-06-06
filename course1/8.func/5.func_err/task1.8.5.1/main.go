package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	initValue, finalValue := "10.0", "15.0"
	percentChange, err := CalculatePercentageChange(initValue, finalValue)
	if err != nil {
		panic(err)
	}
	fmt.Printf("init value: %s, final value: %s, percent change: %.2f%%\n",
		initValue, finalValue, percentChange)
}

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	initialValueFloat, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, err
	}
	finalValueFloat, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}
	change := (math.Abs(finalValueFloat-initialValueFloat) / initialValueFloat) * 100
	return math.Round(change*100) / 100, nil
}
