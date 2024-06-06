package main

import "fmt"

func Sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func Mul(a ...int) int {
	product := 1
	for _, v := range a {
		product *= v
	}
	return product
}

func Sub(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	result := a[0]
	for i := 1; i < len(a); i++ {
		result -= a[i]
	}
	return result
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Mul, 1, 7, 3))
	fmt.Println(MathOperate(Sub, 13, 2, 3))
}
