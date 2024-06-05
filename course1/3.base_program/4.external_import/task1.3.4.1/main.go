package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	var a string = "5.0"
	var b string = "5.0"
	var prec int = 2

	decSum, err := DecimalSum(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Decimal Sum of %s and %s is: %s\n", a, b, decSum)
	}

	decSub, err := DecimalSubtract(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Decimal Subtract of %s and %s is: %s\n", a, b, decSub)
	}

	decMul, err := DecimalMultiply(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Decimal Multiply of %s and %s is: %s\n", a, b, decMul)
	}

	decDiv, err := DecimalDivide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Decimal Divide of %s and %s is: %s\n", a, b, decDiv)
	}

	decRound, err := DecimalRound(a, int32(prec))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Decimal Round of %s with precision %s is: %s\n", a, prec, decRound)
	}

	decGreater, err := DecimalGreaterThan(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		if decGreater {
			fmt.Printf("Decimal %s is greater than %s\n", a, b)
		} else {
			fmt.Printf("Decimal %s is NOT greater than %s\n", b, a)
		}
	}

	decLess, err := DecimalLessThan(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		if decLess {
			fmt.Printf("Decimal %s is less than %s\n", a, b)
		} else {
			fmt.Printf("Decimal %s is NOT less than %s\n", b, a)
		}
	}

	decEqual, err := DecimalEqual(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		if decEqual {
			fmt.Printf("Decimal %s is equal to %s\n", a, b)
		} else {
			fmt.Printf("Decimal %s is NOT equal to %s\n", a, b)
		}
	}
}

func DecimalSum(a, b string) (string, error) {
	aDec, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	bDec, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	sum := decimal.Sum(aDec, bDec)

	return sum.String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	a_dec, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	b_dec, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	sub := a_dec.Sub(b_dec)

	return sub.String(), nil
}

func DecimalMultiply(a, b string) (string, error) {
	a_dec, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	b_dec, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	mul := a_dec.Mul(b_dec)

	return mul.String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	a_dec, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	b_dec, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	div := a_dec.Div(b_dec)

	return div.String(), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	a_dec, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	rounded := a_dec.Round(precision)

	return rounded.String(), nil
}

func DecimalGreaterThan(a, b string) (bool, error) {
	aDec, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}
	bDec, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	if aDec.GreaterThan(bDec) {
		return false, nil
	}
	return true, nil
}

func DecimalLessThan(a, b string) (bool, error) {
	aDec, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}
	bDec, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	if !aDec.LessThan(bDec) {
		return false, nil
	}
	return true, nil
}

func DecimalEqual(a, b string) (bool, error) {
	aDec, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}
	bDec, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	if !aDec.Equal(bDec) {
		return false, nil
	}
	return true, nil
}
