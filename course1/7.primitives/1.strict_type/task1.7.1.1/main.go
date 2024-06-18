package main

import "fmt"

func main() {
	var name string
	var age int
	var city string

	fmt.Print("Введите ваше имя: ")
	_, _ = fmt.Scanln(&name)

	fmt.Print("Введите ваш возраст: ")
	_, _ = fmt.Scanln(&age)

	fmt.Print("Введите ваш город: ")
	_, _ = fmt.Scanln(&city)

	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Город: %s\n", city)
}
