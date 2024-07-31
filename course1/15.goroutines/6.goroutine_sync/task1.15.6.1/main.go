package main

import (
	"fmt"
	"strings"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result strings.Builder

	for _, goroutine := range goroutines {
		wg.Add(1)
		go func(goroutine string) {
			defer wg.Done()
			mu.Lock()
			result.WriteString(goroutine + "\n")
			mu.Unlock()
		}(goroutine())
	}

	wg.Wait()

	return result.String()
}

func main() {
	count := 1000
	goroutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		j := i
		goroutines[j] = func() string {
			return fmt.Sprintf("goroutine %d done", j)
		}
	}

	fmt.Println(waitGroupExample(goroutines...))
}
