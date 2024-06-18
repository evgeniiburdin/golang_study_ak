package main

import "fmt"

func main() {
	fmt.Println(UserInfo("Jane", "Los Angeles", "987-654-3210", 25, 150))
}

func UserInfo(name, city, phone string, age, weight int) string {
	return fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d, Вес: %d",
		name, city, phone, age, weight)
}
