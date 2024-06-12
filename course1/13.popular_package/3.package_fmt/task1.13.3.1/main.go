package main

import (
	"fmt"
	"strconv"
)

func generateMathString(operands []int, operator string) string {
	var result int = operands[0]
	var str string = fmt.Sprintf("%d %s ", operands[0], operator)
	for i := 1; i < len(operands); i++ {
		str += fmt.Sprintf("%d %s ", operands[i], operator)
		switch operator {
		case "+":
			result += operands[i]
		case "-":
			result -= operands[i]
		}
	}
	str = str[:len(str)-3]
	str += " = " + strconv.Itoa(result)
	return str
}
