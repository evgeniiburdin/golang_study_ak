package main

import (
	"fmt"
)

// Parking представляет собой структуру, которая содержит список припаркованных автомобилей
type Parking struct {
	cars []string
}

// Park добавляет автомобиль на парковку
func (p *Parking) Park(carNumber string) {
	p.cars = append(p.cars, carNumber)
	fmt.Printf("Автомобиль %s припаркован.\n", carNumber)
}

// Leave удаляет последний припаркованный автомобиль с парковки
func (p *Parking) Leave() {
	if len(p.cars) == 0 {
		fmt.Println("Парковка пуста.")
		return
	}
	lastIndex := len(p.cars) - 1
	carNumber := p.cars[lastIndex]
	p.cars = p.cars[:lastIndex]
	fmt.Printf("Автомобиль %s покинул парковку.\n", carNumber)
}

func main() {
	parking := &Parking{}

	// Парковка автомобилей
	parking.Park("XYZ-123")
	parking.Park("XYZ-456")
	parking.Park("XYZ-789")

	// Автомобили покидают парковку
	parking.Leave()
	parking.Leave()
	parking.Leave()
	parking.Leave()
}
