package main

import (
	"fmt"
	"strings"

	"strconv"
)

func generateMathString(operands []int, operator string) string {
	var result int = operands[0]
	var str strings.Builder
	str.WriteString(fmt.Sprintf("%d %s ", operands[0], operator))
	for i := 1; i < len(operands); i++ {
		str.WriteString(fmt.Sprintf("%d %s ", operands[i], operator))
		switch operator {
		case "+":
			result += operands[i]
		case "-":
			result -= operands[i]
		}
	}
	return str.String()[:len(str.String())-3] + " = " + strconv.Itoa(result)
}
