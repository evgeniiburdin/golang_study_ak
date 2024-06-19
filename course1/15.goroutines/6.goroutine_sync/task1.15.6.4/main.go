package main

import (
	"fmt"

	"strconv"
	"strings"

	"sync"

	"time"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewCache() *Cache {
	return &Cache{
		users: make(map[string]*User),
	}
}

func (c *Cache) Set(key string, user *User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.users[key] = user
}

func (c *Cache) Get(key string) *User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.users[key]
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprintf("user-%d", i),
		})
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(i))))
		}(i)
	}

	time.Sleep(1 * time.Second)
}
