package main

import (
	"fmt"

	"time"
)

func main() {
	// создаем новый тикер с интервалом 1 секунд
	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, msg string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		timeout := time.After(d)

		for {
			select {
			case <-timeout:
				return
			case <-ticker.C:
				ch <- msg
			}
		}
	}()

	return ch
}
