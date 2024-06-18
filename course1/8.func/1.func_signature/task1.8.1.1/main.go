package main

import (
	"fmt"

	"math"
)

var CalculateCircleArea func(radius float64) float64
var CalculateRectangleArea func(width, height float64) float64
var CalculateTriangleArea func(base, height float64) float64

func main() {
	CalculateCircleArea = func(radius float64) float64 {
		return math.Pi * math.Pow(radius, 2)
	}

	CalculateRectangleArea = func(width, height float64) float64 {
		return width * height
	}

	CalculateTriangleArea = func(base, height float64) float64 {
		return base / 2 * height
	}

	fmt.Printf("Circle area with radius 3 is: %f\n", CalculateCircleArea(3))
	fmt.Printf("Rectangle area with width 3 and height 4 is: %f\n", CalculateRectangleArea(3, 4))
	fmt.Printf("Triangle area with base 3 and height 4 is: %f\n", CalculateTriangleArea(3, 4))
}
