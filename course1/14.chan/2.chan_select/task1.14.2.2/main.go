package main

import (
	"fmt"
	"time"
)

func timeout(timeout time.Duration) func() bool {
	timeoutCh := make(chan bool, 1)

	go func() {
		time.Sleep(timeout)
		timeoutCh <- true
		close(timeoutCh)
	}()

	return func() bool {
		select {
		case <-timeoutCh:
			return false
		default:
			return true
		}
	}
}

func main() {
	timeoutFunc := timeout(1 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)

	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}
