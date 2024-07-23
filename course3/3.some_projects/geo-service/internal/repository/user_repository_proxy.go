package repository

import (
	"errors"
	"log"

	"github.com/go-redis/redis"
)

// RedisCache структура для работы с Redis
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache создает новый экземпляр RedisCache
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

// Set устанавливает значение в кэш
func (c *RedisCache) Set(key string, value interface{}) error {
	err := c.client.Set(key, value, 0).Err()
	if errors.Is(err, redis.Nil) {
		return nil
	}
	return err
}

// Get получает значение из кэша
func (c *RedisCache) Get(key string) (interface{}, error) {
	val, err := c.client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

// Unset удаляет значение из кэша
func (c *RedisCache) Unset(key string) error {
	err := c.client.Del(key).Err()
	if err == redis.Nil {
		return nil
	}
	return err
}

// RedisUserRepositoryProxy структура прокси-репозитория пользователей с кэшированием
type RedisUserRepositoryProxy struct {
	UserRepository UserRepository
	Cache          Cache
}

// NewRedisUserRepositoryProxy создает новый экземпляр RedisUserRepositoryProxy
func NewRedisUserRepositoryProxy(ur UserRepository, r Cache) UserRepositoryProxy {
	return &RedisUserRepositoryProxy{
		UserRepository: ur,
		Cache:          r,
	}
}

// CreateUser создает пользователя
func (ur *RedisUserRepositoryProxy) CreateUser(username, password string) error {
	pass, err := ur.Cache.Get(username)
	if err != nil {
		return err
	}
	if pass != "" {
		return errors.New("user already exists")
	}

	pass, err = ur.UserRepository.ReadUser(username)
	if err != nil && !errors.Is(err, errors.New("user not found")) {
		return err
	}
	if pass != "" {
		return errors.New("user already exists")
	}

	err = ur.UserRepository.CreateUser(username, password)
	if err != nil {
		return err
	}

	return ur.Cache.Set(username, password)
}

// ReadUser читает пользователя
func (ur *RedisUserRepositoryProxy) ReadUser(username string) (password string, err error) {
	pass, err := ur.Cache.Get(username)
	if err != nil {
		return "", err
	}
	if pass != "" {
		log.Println("got user from cache")
		return pass.(string), nil
	}

	pass, err = ur.UserRepository.ReadUser(username)
	if err != nil {
		return "", err
	}
	if pass != "" {
		log.Println("got user from db")
		err = ur.Cache.Set(username, pass)
		if err != nil {
			return "", err
		}
		return pass.(string), nil
	}

	return "", errors.New("user not found")
}

// UpdateUser обновляет пользователя
func (ur *RedisUserRepositoryProxy) UpdateUser(username, newPassword string) error {
	err := ur.UserRepository.UpdateUser(username, newPassword)
	if err != nil {
		return err
	}

	return ur.Cache.Set(username, newPassword)
}

// DeleteUser удаляет пользователя
func (ur *RedisUserRepositoryProxy) DeleteUser(username string) error {
	err := ur.UserRepository.DeleteUser(username)
	if err != nil {
		return err
	}

	return ur.Cache.Unset(username)
}
