package main

import (
	"fmt"
	"sync"
)

func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup

	for _, ch := range from {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for val := range ch {
				mergeTo <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergeTo)
	}()
}

func mergeChan2(chans ...chan int) chan int {
	mainCh := make(chan int)

	wg := sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for val := range c {
				mainCh <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mainCh)
	}()

	return mainCh
}

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	mainCh := make(chan int, 6)
	chs := make([]chan int, 0, 2)

	for i := 0; i < 4; i++ {
		chs = append(chs, generateChan(i))
	}
	mergeChan(mainCh, chs...)
	for val := range mainCh {
		fmt.Printf("%d ", val)
	}

	fmt.Printf("\nmergeChan result\n\n")

	for i := 0; i < 4; i++ {
		chs = append(chs, generateChan(i))
	}
	ch2 := mergeChan2(chs...)
	for val := range ch2 {
		fmt.Printf("%d ", val)
	}

	fmt.Printf("\nmergeChan2 result\n\n")

}
