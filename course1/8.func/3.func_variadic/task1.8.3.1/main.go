package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(UserInfo("John", 21, "Moscow", "Saint Petersburg"))
}

func UserInfo(name string, age int, cities ...string) string {
	return fmt.Sprintf("Имя: %s, возраст: %d, города: %s", name, age, strings.Join(cities, ", "))
}
