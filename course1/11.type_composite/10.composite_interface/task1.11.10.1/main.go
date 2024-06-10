package main

import "fmt"

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case nil:
		return "Пустой интерфейс"
	case []int:

		return "[]int"
	case []float64:

		return "[]float64"
	case []string:

		return "[]string"
	case map[string]interface{}:
		return "map[string]interface{}"
	default:
		return fmt.Sprintf("unknown type: %T", i)
	}
}
