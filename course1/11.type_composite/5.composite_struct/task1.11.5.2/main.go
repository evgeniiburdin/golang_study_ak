package main

import (
	"fmt"

	"github.com/icrowley/fake"
	"math/rand"
)

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	animals := make([]Animal, 0, 3)
	for i := 0; i < 3; i++ {
		animals = append(animals, Animal{
			Type: "animal",
			Name: fake.FirstName(),
			Age:  rand.Intn(20),
		})
	}
	return animals
}

func preparePrint(animals []Animal) string {
	str := ""
	for _, animal := range animals {
		str += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age)
	}
	return str
}
