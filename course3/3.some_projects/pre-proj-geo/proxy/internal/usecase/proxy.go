package usecase

import (
	"context"
	"geo-service-proxy/internal/entity"
	pb_auth "geo-service-proxy/internal/usecase/auth-service/auth"
	pb_geo "geo-service-proxy/internal/usecase/geo-service/geo"
)

// ProxyUseCase -.
type ProxyUseCase struct {
	authServiceClient pb_auth.AuthServiceClient
	geoServiceClient  pb_geo.GeoServiceClient
}

// New -.
func New(asc pb_auth.AuthServiceClient, gsc pb_geo.GeoServiceClient) *ProxyUseCase {
	return &ProxyUseCase{
		authServiceClient: asc,
		geoServiceClient:  gsc,
	}
}

func (uc *ProxyUseCase) CreateUser(ctx context.Context, user entity.User) (accessToken, refreshToken string, err error) {
	resp, err := uc.authServiceClient.CreateUser(ctx, &pb_auth.CreateUserRequest{
		User: &pb_auth.User{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
	if err != nil {
		return "", "", err
	}

	return resp.AccessToken, resp.RefreshToken, nil
}

func (uc *ProxyUseCase) GetUser(ctx context.Context, accessToken string) (entity.User, error) {
	resp, err := uc.authServiceClient.GetUser(ctx, &pb_auth.GetUserRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Username: resp.User.Username,
		Password: resp.User.Password,
		Email:    resp.User.Email,
		Role:     resp.User.Role,
	}

	return user, nil
}

func (uc *ProxyUseCase) GetUsers(ctx context.Context, accessToken string) ([]entity.User, error) {
	resp, err := uc.authServiceClient.GetUsers(ctx, &pb_auth.GetUsersRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return []entity.User{}, err
	}

	users := make([]entity.User, len(resp.Users))
	for i, u := range resp.Users {
		users[i] = entity.User{
			Username: u.Username,
			Password: u.Password,
			Email:    u.Email,
			Role:     u.Role,
		}
	}

	return users, nil
}

func (uc *ProxyUseCase) UpdateUser(ctx context.Context, accessToken string, user entity.User) error {
	_, err := uc.authServiceClient.UpdateUser(ctx, &pb_auth.UpdateUserRequest{
		AccessToken: accessToken,
		User: &pb_auth.User{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Role:     user.Role,
		},
	})

	return err
}

func (uc *ProxyUseCase) DeleteUser(ctx context.Context, accessToken string) error {
	_, err := uc.authServiceClient.DeleteUser(ctx, &pb_auth.DeleteUserRequest{
		AccessToken: accessToken,
	})

	return err
}

func (uc *ProxyUseCase) Login(ctx context.Context, email, password string) (accessToken, refreshToken string, err error) {
	resp, err := uc.authServiceClient.Login(ctx, &pb_auth.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", "", err
	}

	accessToken = resp.AccessToken
	refreshToken = resp.RefreshToken

	return accessToken, refreshToken, nil
}

func (uc *ProxyUseCase) Logout(ctx context.Context, accessToken string) error {
	_, err := uc.authServiceClient.Logout(ctx, &pb_auth.LogoutRequest{
		AccessToken: accessToken,
	})

	return err
}

func (uc *ProxyUseCase) ValidateToken(ctx context.Context, jwt string) error {
	_, err := uc.authServiceClient.ValidateToken(ctx, &pb_auth.ValidateTokenRequest{
		AccessToken: jwt,
	})

	return err
}

func (uc *ProxyUseCase) RefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	resp, err := uc.authServiceClient.RefreshToken(ctx, &pb_auth.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return "", "", err
	}

	newAccessToken = resp.AccessToken
	newRefreshToken = resp.RefreshToken

	return newAccessToken, newRefreshToken, nil
}

func (uc *ProxyUseCase) GeocodeToAddress(ctx context.Context, geocode entity.Geocode) (address entity.Address, err error) {
	resp, err := uc.geoServiceClient.GeocodeToAddress(ctx, &pb_geo.Geocode{
		Lat: geocode.Lat,
		Lng: geocode.Lng,
	})
	if err != nil {
		return address, err
	}

	address = entity.Address{
		Country: resp.Country,
		City:    resp.City,
	}

	return address, nil
}
