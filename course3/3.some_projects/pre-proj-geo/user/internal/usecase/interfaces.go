package usecase

import (
	"context"

	"user-service/internal/entity"
)

type (
	// Userer -.
	Userer interface {
		Write(ctx context.Context, u entity.User) error
		Get(ctx context.Context, email string) (entity.User, error)
		GetAll(ctx context.Context) ([]entity.User, error)
		Update(ctx context.Context, u entity.User) error
		Delete(ctx context.Context, email string) error
	}

	// UserRepo -.
	UserRepo interface {
		Write(ctx context.Context, u entity.User) error
		Get(ctx context.Context, email string) (entity.User, error)
		GetAll(ctx context.Context) ([]entity.User, error)
		Update(ctx context.Context, u entity.User) error
		Delete(ctx context.Context, email string) error
	}

	// UserCache -.
	UserCache interface {
		Set(u entity.User) error
		Get(email string) (entity.User, error)
		Delete(email string) error
	}
)
