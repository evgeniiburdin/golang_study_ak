package main

import "fmt"

// CircuitRinger интерфейс для кольцевого буфера
type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

// RingBuffer структура кольцевого буфера
type RingBuffer struct {
	buffer []int
	size   int
	start  int
	end    int
	count  int
}

// NewRingBuffer создает новый кольцевой буфер заданного размера
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, size),
		size:   size,
		start:  0,
		end:    0,
		count:  0,
	}
}

// Add добавляет значение в буфер
func (rb *RingBuffer) Add(val int) {
	rb.buffer[rb.end] = val
	rb.end = (rb.end + 1) % rb.size
	if rb.count == rb.size {
		rb.start = (rb.start + 1) % rb.size // перезаписывает старое значение
	} else {
		rb.count++
	}
}

// Get возвращает значение из буфера
func (rb *RingBuffer) Get() (int, bool) {
	if rb.count == 0 {
		return 0, false // буфер пуст
	}
	val := rb.buffer[rb.start]
	rb.start = (rb.start + 1) % rb.size
	rb.count--
	return val, true
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4) // Перезаписывает значение 1

	for val, ok := rb.Get(); ok; val, ok = rb.Get() {
		fmt.Println(val) // Выводит: 2, 3, 4
	}

	if _, ok := rb.Get(); !ok {
		fmt.Println("Buffer is empty") // Выводит: Buffer is empty
	}
}
