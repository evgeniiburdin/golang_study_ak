package main

import (
	"fmt"
)

func getVariableType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
