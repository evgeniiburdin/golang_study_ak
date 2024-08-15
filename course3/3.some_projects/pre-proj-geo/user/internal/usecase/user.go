package usecase

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"user-service/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	cache UserCache
	repo  UserRepo
}

// New -.
func New(r UserRepo, c UserCache) *UserUseCase {
	return &UserUseCase{
		repo:  r,
		cache: c,
	}
}

func (u *UserUseCase) Write(ctx context.Context, user entity.User) error {
	/*var user entity.User
	if err := json.Unmarshal(jsonUser, &user); err != nil {
		return fmt.Errorf("UserUseCase - Write - failed to unmarshal user: %w", err)
	}*/

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("UserUseCase - Write: %w", err)
	}
	user.Password = string(encryptedPassword)

	err = u.repo.Write(ctx, user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Write - s.repo.Write: %w", err)
	}

	err = u.cache.Set(user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Write - s.cache.Set: %w", err)
	}

	return nil
}

func (u *UserUseCase) Get(ctx context.Context, email string) (entity.User, error) {
	user, err := u.cache.Get(email)
	if err != nil {
		user, err = u.repo.Get(ctx, email)
		if err != nil {
			return entity.User{}, fmt.Errorf("UserUseCase - Get - s.repo.Get: %w", err)
		}

		return user, nil
	}

	return user, nil
}

func (u *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {
	users, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetAll - s.repo.GetAll: %w", err)
	}

	return users, nil
}

func (u *UserUseCase) Update(ctx context.Context, user entity.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("UserUseCase - Update: %w", err)
	}
	user.Password = string(encryptedPassword)

	err = u.repo.Update(ctx, user)
	if err != nil {
		fmt.Errorf("UserUseCase - Update - s.repo.Update: %w", err)
	}

	err = u.cache.Set(user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Update - s.cache.Set: %w", err)
	}

	return nil
}

func (u *UserUseCase) Delete(ctx context.Context, email string) error {
	err := u.repo.Delete(ctx, email)
	if err != nil {
		return fmt.Errorf("UserUseCase - Delete - s.repo.Delete: %w", err)
	}

	err = u.cache.Delete(email)
	if err != nil {
		return fmt.Errorf("UserUseCase - Delete - s.cache.Delete: %w", err)
	}

	return nil
}
