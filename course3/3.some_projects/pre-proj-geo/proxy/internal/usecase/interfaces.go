package usecase

import (
	"context"

	"geo-service-proxy/internal/entity"
)

type (
	// Proxyer -.
	Proxyer interface {
		CreateUser(ctx context.Context, user entity.User) (accessToken, refreshToken string, err error)
		GetUser(ctx context.Context, accessToken string) (entity.User, error)
		GetUsers(ctx context.Context, accessToken string) ([]entity.User, error)
		UpdateUser(ctx context.Context, accessToken string, user entity.User) error
		DeleteUser(ctx context.Context, accessToken string) error
		Login(ctx context.Context, email, password string) (accessToken, refreshToken string, err error)
		Logout(ctx context.Context, accessToken string) error
		ValidateToken(ctx context.Context, jwt string) error
		RefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error)
		GeocodeToAddress(ctx context.Context, geocode entity.Geocode) (address entity.Address, err error)
	}
)
