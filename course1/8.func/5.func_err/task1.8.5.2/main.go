package main

import (
	"errors"
	"fmt"
)

func main() {
	disc1, err := CheckDiscount(1488, 0.64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(disc1)
	}
	disc2, err := CheckDiscount(1990, 0.48)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(disc2)
	}
}

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 0.5 {
		return 0, errors.New("скидка не может превышать 50%")
	}
	return price * discount, nil
}
