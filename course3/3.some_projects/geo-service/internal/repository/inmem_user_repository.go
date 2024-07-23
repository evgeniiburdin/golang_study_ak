package repository

import (
	"errors"
)

type InMemoryUserRepository struct {
	users map[string]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]string),
	}
}

func (r *InMemoryUserRepository) ReadUser(username string) (string, error) {
	password, ok := r.users[username]
	if !ok {
		return "", errors.New("user not found")
	}
	return password, nil
}

func (r *InMemoryUserRepository) CreateUser(username, password string) error {
	if _, ok := r.users[username]; ok {
		return errors.New("user already exists")
	}
	r.users[username] = password
	return nil
}

func (r *InMemoryUserRepository) UpdateUser(username, newPassword string) error {
	if _, ok := r.users[username]; !ok {
		return errors.New("user not found")
	}
	r.users[username] = newPassword
	return nil
}

func (r *InMemoryUserRepository) DeleteUser(username string) error {
	if _, ok := r.users[username]; !ok {
		return errors.New("user not found")
	}
	delete(r.users, username)
	return nil
}
