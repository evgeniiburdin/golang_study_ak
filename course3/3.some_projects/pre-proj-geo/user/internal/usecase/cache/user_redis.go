package cache

import (
	"encoding/json"
	"errors"

	"user-service/internal/entity"
	redispkg "user-service/pkg/redis"
)

// UserCache -.
type UserCache struct {
	Redis *redispkg.Redis
}

// New -.
func New(r *redispkg.Redis) (*UserCache, error) {
	if r == nil {
		return nil, errors.New("redis client is nil")
	}

	return &UserCache{r}, nil
}

// Write -.
func (c *UserCache) Set(u entity.User) error {
	jsonUser, err := json.Marshal(u)
	if err != nil {
		return err
	}

	err = c.Redis.Client.Set(u.Email, jsonUser, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// Get -.
func (c *UserCache) Get(email string) (entity.User, error) {
	var user entity.User
	err := c.Redis.Client.Get(email).Scan(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// Delete -.
func (c *UserCache) Delete(email string) error {
	err := c.Redis.Client.Del(email).Err()
	if err != nil {
		return err
	}

	return nil
}
