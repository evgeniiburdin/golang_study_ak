package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
)

type Cacher interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type cache struct {
	client *redis.Client
}

func (c *cache) Set(key string, value interface{}) error {
	err := c.client.Set(key, value, 0).Err()
	if errors.Is(err, redis.Nil) {
		return nil
	}
	return err
}

func (c *cache) Get(key string) (interface{}, error) {
	result, err := c.client.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return result, nil
	}
	return nil, err
}

func NewCache(client *redis.Client) Cacher {
	return &cache{
		client: client,
	}
}

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	cache := NewCache(client)

	err := cache.Set("some key", "some value")
	if err != nil {
		panic(err)
	}

	val, err := cache.Get("some key")
	if err != nil {
		panic(err)
	}

	fmt.Println("val:", val)

	user := &User{
		ID:   1,
		Name: "John",
		Age:  30,
	}

	err = cache.Set(fmt.Sprintf("user%d", user.ID), user)
	if err != nil {
		panic(err)
	}

	val, err = cache.Get(fmt.Sprintf("user%d", user.ID))
	if err != nil {
		panic(err)
	}
}
