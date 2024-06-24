package main

import (
	"testing"
)

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Без использования пула создаем новый объект
		p = &Person{Age: i}
		_ = p // Используем объект
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Получаем объект из пула
		p = personPool.Get().(*Person)
		// Инициализируем объект
		p.Age = i
		_ = p // Используем объект
		// Возвращаем объект обратно в пул
		personPool.Put(p)
	}
}
