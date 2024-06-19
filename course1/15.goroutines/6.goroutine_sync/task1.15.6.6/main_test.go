package main

import (
	"sync"

	"testing"
)

func TestCache(t *testing.T) {
	cache := &Cache{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(string(rune(i)), i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, ok := cache.Get(string(rune(i)))
			if !ok || value != i {
				t.Errorf("Expected %d, got %v", i, value)
			}
		}(i)
	}

	wg.Wait()
}
