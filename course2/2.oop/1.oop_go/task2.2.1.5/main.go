package main

import "fmt"

// Mover интерфейс определяет методы для перемещения и получения скоростей
type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

// BaseMover структура содержит базовую скорость
type BaseMover struct {
	speed int
}

// Speed метод возвращает текущую скорость
func (bm BaseMover) Speed() int {
	return bm.speed
}

// MaxSpeed метод возвращает максимальную скорость
func (bm BaseMover) MaxSpeed() int {
	return 120
}

// MinSpeed метод возвращает минимальную скорость
func (bm BaseMover) MinSpeed() int {
	return 10
}

// FastMover структура представляет быстрый объект
type FastMover struct {
	BaseMover
}

// Move метод для FastMover
func (fm FastMover) Move() string {
	return fmt.Sprintf("Fast mover! Moving at speed: %d", fm.Speed())
}

// SlowMover структура представляет медленный объект
type SlowMover struct {
	BaseMover
}

// Move метод для SlowMover
func (sm SlowMover) Move() string {
	return fmt.Sprintf("Slow mover... Moving at speed: %d", sm.Speed())
}

func main() {
	var movers []Mover

	fm := FastMover{BaseMover{100}}
	sm := SlowMover{BaseMover{10}}

	movers = append(movers, fm, sm)

	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}
