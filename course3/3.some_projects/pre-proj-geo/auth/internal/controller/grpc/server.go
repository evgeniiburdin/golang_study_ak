package grpc

import (
	"auth-service/api/auth/gen/auth-service/auth"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"auth-service/internal/entity"
	"auth-service/internal/usecase"
	"auth-service/pkg/logger"
)

const (
	defaultShutdownTimeout = 3 * time.Second
)

type GRPCServer struct {
	auth.UnimplementedAuthServiceServer
	server          *grpc.Server
	listener        net.Listener
	notify          chan error
	shutdownTimeout time.Duration
	uc              usecase.Auther
	lg              logger.Interface
}

func NewGRPCServer(uc usecase.Auther, lg logger.Interface, opts ...Option) *GRPCServer {
	grpcServer := &GRPCServer{
		server:          grpc.NewServer(),
		listener:        nil,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
		uc:              uc,
		lg:              lg,
	}
	auth.RegisterAuthServiceServer(grpcServer.server, grpcServer)

	// Custom options
	for _, opt := range opts {
		opt(grpcServer)
	}

	grpcServer.start()

	return grpcServer
}

func (s *GRPCServer) start() {
	go func() {
		s.notify <- s.server.Serve(s.listener)
		close(s.notify)
	}()
}

// Notify -.
func (s *GRPCServer) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *GRPCServer) Shutdown() {
	s.server.GracefulStop()
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	startTime := time.Now()

	accessToken, refreshToken, err := s.uc.CreateUser(ctx, entity.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password,
		Role:     req.User.Role,
	})
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - CreateUser: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to write user: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - CreateUser: request { %#v } completed in %dms with response { accessToken: %s, refreshToken: %s, error: %#v }",
			req, timeTaken.Milliseconds(), accessToken, refreshToken, err))
	}()

	return &auth.CreateUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *GRPCServer) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	startTime := time.Now()

	user, err := s.uc.GetUser(ctx, req.AccessToken)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - GetUser: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to read user: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - GetUser: request { %#v } completed in %dms with response { user: %#v, error: %#v }",
			req, timeTaken.Milliseconds(), user, err))
	}()

	return &auth.GetUserResponse{
		User: &auth.User{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	}, nil
}

func (s *GRPCServer) GetUsers(ctx context.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	startTime := time.Now()

	users, err := s.uc.GetUsers(ctx, req.AccessToken)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - GetUsers: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to read users: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - GetUsers: request { %#v } completed in %dms with response { users: %+v, error: %#v }",
			req, timeTaken.Milliseconds(), users, err))
	}()

	respUsers := make([]*auth.User, len(users))
	for i, user := range users {
		respUsers[i] = &auth.User{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		}
	}

	return &auth.GetUsersResponse{
		Users: respUsers,
	}, nil
}

func (s *GRPCServer) UpdateUser(ctx context.Context, req *auth.UpdateUserRequest) (*emptypb.Empty, error) {
	startTime := time.Now()

	err := s.uc.UpdateUser(ctx, req.AccessToken, entity.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Role:     req.User.Role,
		Password: req.User.Password,
	})
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - UpdateUser: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - UpdateUser: request { %#v  %#v } completed in %dms with response { error: %#v }",
			req.User, req.AccessToken, timeTaken.Milliseconds(), err))
	}()

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) (*emptypb.Empty, error) {
	startTime := time.Now()

	err := s.uc.DeleteUser(ctx, req.AccessToken)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - DeleteUser: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - DeleteUser: request { %#v } completed in %dms with response { error: %#v }",
			req, timeTaken.Milliseconds(), err))
	}()

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	startTime := time.Now()

	accessToken, refreshToken, err := s.uc.Login(ctx, req.Email, req.Password)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - Login: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to login: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - Login: request { %#v } completed in %dms with response { accessToken: %s, refreshToken: %s, error: %#v }",
			req, timeTaken.Milliseconds(), accessToken, refreshToken, err))
	}()

	return &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *GRPCServer) Logout(ctx context.Context, req *auth.LogoutRequest) (*emptypb.Empty, error) {
	startTime := time.Now()

	err := s.uc.Logout(ctx, req.AccessToken)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - Logout: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to logout: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - Logout: request { %#v } completed in %dms with response { error: %#v }",
			req, timeTaken.Milliseconds(), err))
	}()

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*emptypb.Empty, error) {
	startTime := time.Now()

	err := s.uc.ValidateToken(ctx, req.AccessToken)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - ValidateToken: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to validate token: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - ValidateToken: request { %#v } completed in %dms with response { error: %#v }",
			req, timeTaken.Milliseconds(), err))
	}()

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	startTime := time.Now()

	accessToken, refreshToken, err := s.uc.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - RefreshToken: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to refresh token: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - RefreshToken: request { %#v } completed in %dms with response { accessToken: %s, refreshToken: %s, error: %#v }",
			req, timeTaken.Milliseconds(), accessToken, refreshToken, err))
	}()

	return &auth.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
