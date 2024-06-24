package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// User struct implements "user" object with two fields:
// ID (type int),
// Name (type string)
type User struct {
	ID   int
	Name string
}

// Cache struct implements "cache" object with two fields:
// mu (type sync.RWMutex) for concurrent operations security
// items (type map[string]interface{}) as a storage for objects.
type Cache struct {
	mu    sync.RWMutex
	items map[string]interface{}
}

// NewCache generates and return a new Cache object with initialized `items` map
func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

// Set adds an object to the cache using the specified key. Uses Lock to provide concurrent write security.
func (c *Cache) Set(key string, item interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = item
}

// Get returns an object from cache of the specified key. Uses RLock to provide concurrent read security.
func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.items[key]
}

// keyBuilder concatenates strings given as args, separating them with `:`
func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

// GetUser converts interface{} type to *User type. The object retrieved from the cache is expected to be of type *User.
func GetUser(i interface{}) *User {
	return i.(*User)
}

// main creates a new Cache object, runs 100 goroutines adding 100 `User` objs to the Cache, runs 100 goroutines
// retrieving objects from the Cache and printing them out
func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprintf("user-%d", i),
		})
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}

	time.Sleep(1 * time.Second)
}
