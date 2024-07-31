package main

import (
	"fmt"
	"hash/crc32"
	"sync"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HashMap struct {
	mu       sync.Mutex
	data     map[uint32]interface{}
	hashFunc func(string) uint32
}

func NewHashMap(options ...func(*HashMap)) *HashMap {
	hm := &HashMap{
		data:     make(map[uint32]interface{}),
		hashFunc: defaultHashFunc,
	}
	for _, option := range options {
		option(hm)
	}
	return hm
}

func defaultHashFunc(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func WithHashFunc(hashFunc func(string) uint32) func(*HashMap) {
	return func(hm *HashMap) {
		hm.hashFunc = hashFunc
	}
}

func (hm *HashMap) Set(key string, value interface{}) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hashedKey := hm.hashFunc(key)
	hm.data[hashedKey] = value
}

func (hm *HashMap) Get(key string) (interface{}, bool) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hashedKey := hm.hashFunc(key)
	value, ok := hm.data[hashedKey]
	return value, ok
}

func MeasureTime(f func()) time.Duration {
	start := time.Now()
	f()
	since := time.Since(start)
	return since
}

func TestSlice16() {
	m := NewHashMap()
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap\n")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'\n", i, value)
		}
	}
}

func TestSlice1000() {
	m := NewHashMap()
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap\n")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'\n", i, value)
		}
	}
}

func main() {
	time := MeasureTime(TestSlice16)
	fmt.Println(time)
	time = MeasureTime(TestSlice1000)
	fmt.Println(time)
}
