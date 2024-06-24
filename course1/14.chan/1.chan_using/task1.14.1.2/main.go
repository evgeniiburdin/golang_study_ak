package main

import (
	"fmt"
	"time"
)

func generateData(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- 1
		}
		close(out)
	}()
	return out
}

func main() {
	data := generateData(10)

	go func() {
		time.Sleep(1 * time.Second)
	}()

	for v := range data {
		fmt.Println(v)
	}
}
