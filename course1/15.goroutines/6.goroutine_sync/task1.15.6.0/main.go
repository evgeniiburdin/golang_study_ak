package main

import (
	"fmt"

	"sync"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
}

func main() {
	// Получаем объект из пула
	p := personPool.Get().(*Person)
	// Используем объект
	p.Age = 30
	fmt.Printf("Person age: %d\n", p.Age)
	// Возвращаем объект обратно в пул
	personPool.Put(p)
}
