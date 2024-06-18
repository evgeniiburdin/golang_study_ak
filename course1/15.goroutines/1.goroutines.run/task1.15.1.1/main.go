package main

import (
	"fmt"

	"time"
)

func main() {
	stop := make(chan bool, 1)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина завершила работу")
		stop <- true
	}()

	timer := time.NewTimer(5 * time.Second)

	data := NotifyOnTimer(timer, stop)

	for v := range data {
		fmt.Println(v)
	}
}

func NotifyOnTimer(timer *time.Timer, stop chan bool) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		select {
		case <-timer.C:
			ch <- "Таймер сработал раньше, чем горутина завершила работу"
		case <-stop:
			ch <- "Горутина завершила работу раньше, чем таймер сработал"
		}
	}()

	return ch
}
