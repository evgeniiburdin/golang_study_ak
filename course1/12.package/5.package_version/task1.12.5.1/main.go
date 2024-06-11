package main

import (
	"fmt"
	"github.com/evgeniiburdin/mymath"
)

func main() {
	fmt.Println(mymath.Ceil(5.5))
	fmt.Println(mymath.Floor(5.5))
	fmt.Println(mymath.Max(5.0, 4.0))
	fmt.Println(mymath.Min(5.0, 4.0))
	fmt.Println(mymath.Pow(5.0, 2.0))
	fmt.Println(mymath.Sqrt(25.0))
}
