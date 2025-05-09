package usecase

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"

	"auth-service/internal/entity"
	pb "auth-service/internal/usecase/user-service/user"
	jwtpkg "auth-service/pkg/jwt"
)

// AuthUseCase -.
type AuthUseCase struct {
	userServiceClient pb.UserServiceClient
	authRepo          AuthRepo
	jwtService        jwtpkg.JWTServicer
}

// New -.
func New(usc pb.UserServiceClient, r AuthRepo, js jwtpkg.JWTServicer) *AuthUseCase {
	return &AuthUseCase{
		userServiceClient: usc,
		authRepo:          r,
		jwtService:        js,
	}
}

func (u *AuthUseCase) CreateUser(ctx context.Context, user entity.User) (accessToken, refreshToken string, err error) {
	// gRPC data send scenario

	_, err = u.userServiceClient.Write(ctx, &pb.WriteUserRequest{
		User: &pb.User{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - CreateUser - s.userServiceClient: %w", err)
	}

	// Kafka data send scenario
	/*
		err = u.kafkaProducer.SerializeAndProduce(user, "createUser")
		if err != nil {
			return "", "", err
		}
	*/
	accessToken, refreshToken, err = u.jwtService.CreateJWTs(user)
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - CreateUser - s.jwtService: %w", err)
	}

	err = u.authRepo.WriteRefreshToken(ctx, user.Username, refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - CreateUser - u.authRepo.WriteRefreshToken: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (u *AuthUseCase) GetUser(ctx context.Context, accessToken string) (entity.User, error) {
	jwtClaims, err := u.jwtService.ReadJWT(accessToken)
	if err != nil {
		return entity.User{}, fmt.Errorf("AuthUseCase - GetUser - s.jwtService: %w", err)
	}

	userEmail, ok := jwtClaims["sub"]
	if !ok {
		return entity.User{}, fmt.Errorf(`AuthUseCase - GetUser - "sub" field not found in jwtClaims`)
	}

	resp, err := u.userServiceClient.Get(ctx, &pb.GetUserRequest{
		Email: userEmail.(string),
	})
	if err != nil {
		return entity.User{}, fmt.Errorf("AuthUseCase - GetUser - s.userServiceClient: %w", err)
	}

	return entity.User{
		Username: resp.User.Username,
		Password: resp.User.Password,
		Email:    resp.User.Email,
		Role:     resp.User.Role,
	}, nil
}

func (u *AuthUseCase) GetUsers(ctx context.Context, accessToken string) ([]entity.User, error) {
	jwtClaims, err := u.jwtService.ReadJWT(accessToken)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - GetUsers - s.jwtService: %w", err)
	}

	userRole, ok := jwtClaims["role"]
	if !ok {
		return nil, fmt.Errorf(`AuthUseCase - GetUser - "role" field not found in jwtClaims`)
	}

	if userRole.(string) != "admin" {
		return nil, fmt.Errorf(`AuthUseCase - GetUser - "role" field is not "admin"`)
	}

	resp, err := u.userServiceClient.GetAll(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - GetUsers - s.userServiceClient: %w", err)
	}

	users := make([]entity.User, 0, len(resp.Users))
	for _, user := range resp.Users {
		users = append(users, entity.User{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	return users, nil
}

func (u *AuthUseCase) UpdateUser(ctx context.Context, accessToken string, user entity.User) error {
	jwtClaims, err := u.jwtService.ReadJWT(accessToken)
	if err != nil {
		return fmt.Errorf("AuthUseCase - UpdateUser - s.jwtService: %w", err)
	}

	userEmail, ok := jwtClaims["sub"]
	if !ok {
		return fmt.Errorf(`AuthUseCase - GetUser - "sub" field not found in jwtClaims`)
	}

	_, err = u.userServiceClient.Update(ctx, &pb.UpdateUserRequest{
		User: &pb.User{
			Username: user.Username,
			Password: user.Password,
			Email:    userEmail.(string),
			Role:     user.Role,
		},
	})
	if err != nil {
		return fmt.Errorf("AuthUseCase - UpdateUser - s.userServiceClient: %w", err)
	}

	return nil
}

func (u *AuthUseCase) DeleteUser(ctx context.Context, accessToken string) error {
	jwtClaims, err := u.jwtService.ReadJWT(accessToken)
	if err != nil {
		return fmt.Errorf("AuthUseCase - UpdateUser - s.jwtService: %w", err)
	}

	userEmail, ok := jwtClaims["sub"]
	if !ok {
		return fmt.Errorf(`AuthUseCase - GetUser - "sub" field not found in jwtClaims`)
	}

	_, err = u.userServiceClient.Delete(ctx, &pb.DeleteUserRequest{
		Email: userEmail.(string),
	})
	if err != nil {
		return fmt.Errorf("AuthUseCase - UpdateUser - s.userServiceClient: %w", err)
	}

	return nil
}

func (u *AuthUseCase) Login(ctx context.Context, email, password string) (accessToken, refreshToken string, err error) {
	resp, err := u.userServiceClient.Get(ctx, &pb.GetUserRequest{
		Email: email,
	})
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - Login - s.userServiceClient: %w", err)
	}

	if resp.User.Email != email {
		return "", "", fmt.Errorf("AuthUseCase - Login - unauthorized: wrong email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.User.Password), []byte(password))
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - Login - unauthorized: wrong password")
	}

	accessToken, refreshToken, err = u.jwtService.CreateJWTs(entity.User{
		Username: resp.User.Username,
		Password: resp.User.Password,
		Email:    resp.User.Email,
		Role:     resp.User.Role,
	})
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - Login - s.jwtService: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (u *AuthUseCase) Logout(ctx context.Context, accessToken string) error {
	jwtClaims, err := u.jwtService.ReadJWT(accessToken)
	if err != nil {
		return fmt.Errorf("AuthUseCase - UpdateUser - s.jwtService: %w", err)
	}

	userEmail, ok := jwtClaims["sub"]
	if !ok {
		return fmt.Errorf(`AuthUseCase - GetUser - "sub" field not found in jwtClaims`)
	}

	err = u.authRepo.RemoveRefreshToken(ctx, userEmail.(string))
	if err != nil {
		return fmt.Errorf("AuthUseCase - UpdateUser - s.authRepo.RemoveRefreshToken: %w", err)
	}

	return nil
}

func (u *AuthUseCase) ValidateToken(ctx context.Context, accessToken string) error {
	_, err := u.jwtService.ReadJWT(accessToken)
	if err != nil {
		return fmt.Errorf("AuthUseCase - ValidateToken - s.jwtService: %w", err)
	}

	return nil
}

func (u *AuthUseCase) RefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	err = u.authRepo.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - ValidateToken - s.authRepo.ValidateRefreshToken: %w", err)
	}

	jwtClaims, err := u.jwtService.ReadJWT(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - ValidateToken - s.jwtService: %w", err)
	}

	userEmail, ok := jwtClaims["sub"]
	if !ok {
		return "", "", fmt.Errorf(`AuthUseCase - RefreshToken - "sub" field not found in jwtClaims`)
	}

	resp, err := u.userServiceClient.Get(ctx, &pb.GetUserRequest{
		Email: userEmail.(string),
	})
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - ValidateToken - s.userServiceClient: %w", err)
	}

	newAccessToken, newRefreshToken, err = u.jwtService.CreateJWTs(entity.User{
		Username: resp.User.Username,
		Password: resp.User.Password,
		Email:    resp.User.Email,
		Role:     resp.User.Role,
	})
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - ValidateToken - s.jwtService: %w", err)
	}

	err = u.authRepo.WriteRefreshToken(ctx, resp.User.Email, newRefreshToken)
	if err != nil {
		return "", "", fmt.Errorf("AuthUseCase - CreateUser - u.authRepo.WriteRefreshToken: %w", err)
	}

	return newAccessToken, newRefreshToken, nil
}
