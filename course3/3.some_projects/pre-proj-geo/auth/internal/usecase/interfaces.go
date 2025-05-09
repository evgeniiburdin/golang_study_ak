package usecase

import (
	"context"

	"auth-service/internal/entity"
)

type (
	// Auther -.
	Auther interface {
		CreateUser(ctx context.Context, user entity.User) (accessToken, refreshToken string, err error)
		GetUser(ctx context.Context, accessToken string) (entity.User, error)
		GetUsers(ctx context.Context, accessToken string) ([]entity.User, error)
		UpdateUser(ctx context.Context, accessToken string, user entity.User) error
		DeleteUser(ctx context.Context, accessToken string) error
		Login(ctx context.Context, email, password string) (accessToken, refreshToken string, err error)
		Logout(ctx context.Context, accessToken string) error
		ValidateToken(ctx context.Context, accessToken string) error
		RefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error)
	}

	// AuthRepo -.
	AuthRepo interface {
		WriteRefreshToken(ctx context.Context, email, refreshToken string) error
		ValidateRefreshToken(ctx context.Context, refreshToken string) error
		UpdateRefreshToken(ctx context.Context, email, refreshToken string) error
		RemoveRefreshToken(ctx context.Context, email string) error
	}
)
