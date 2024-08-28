package main

import (
	"fmt"
)

// Bank представляет собой структуру, которая содержит очередь клиентов
type Bank struct {
	queue []string
}

// AddClient добавляет клиента в очередь
func (b *Bank) AddClient(client string) {
	b.queue = append(b.queue, client)
}

// ServeNextClient обслуживает следующего клиента в очереди и удаляет его из очереди
func (b *Bank) ServeNextClient() string {
	if len(b.queue) == 0 {
		return "No clients in the queue"
	}
	nextClient := b.queue[0]
	b.queue = b.queue[1:]
	return nextClient
}

func main() {
	bank := Bank{}
	bank.AddClient("Client 1")
	bank.AddClient("Client 2")
	bank.AddClient("Client 3")

	fmt.Println(bank.ServeNextClient()) // Output: Client 1
	fmt.Println(bank.ServeNextClient()) // Output: Client 2
	fmt.Println(bank.ServeNextClient()) // Output: Client 3
	fmt.Println(bank.ServeNextClient()) // Output: No clients in the queue
}
