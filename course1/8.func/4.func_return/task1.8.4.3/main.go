package main

import "fmt"

func main() {
	fmt.Println(CalculateStockValue(45.673, 367))
}

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	return price * float64(quantity), price
}
