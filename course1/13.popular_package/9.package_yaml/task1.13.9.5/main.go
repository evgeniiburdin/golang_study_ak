package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func main() {
	data := []byte(`{"name":"John", "age":30}`)
	var person Person
	err := unmarshal(data, &person)
	if err != nil {
		fmt.Println("Ошибка декодирования данных:", err)
		return
	}
	fmt.Println("Имя: ", person.Name)
	fmt.Println("Возраст: ", person.Age)
}
