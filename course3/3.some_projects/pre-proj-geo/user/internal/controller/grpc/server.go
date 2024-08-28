package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "user-service/internal/controller/grpc/gen/user-service/user"
	"user-service/internal/entity"
	"user-service/internal/usecase"
	"user-service/pkg/logger"
)

const (
	defaultShutdownTimeout = 3 * time.Second
)

type GRPCServer struct {
	pb.UnimplementedUserServiceServer
	server          *grpc.Server
	listener        net.Listener
	notify          chan error
	shutdownTimeout time.Duration
	uc              usecase.Userer
	lg              logger.Interface
}

func NewGRPCServer(uc usecase.Userer, lg logger.Interface, opts ...Option) *GRPCServer {
	grpcServer := &GRPCServer{
		server:          grpc.NewServer(),
		listener:        nil,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
		uc:              uc,
		lg:              lg,
	}
	pb.RegisterUserServiceServer(grpcServer.server, grpcServer)

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

func (s *GRPCServer) Write(ctx context.Context, req *pb.WriteUserRequest) (*emptypb.Empty, error) {
	startTime := time.Now()

	err := s.uc.Write(ctx, entity.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password,
		Role:     req.User.Role,
	})
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - Write: %w", err))
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to write user: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - Write: request { user: %#v } completed in %dms with response { error: %#v }",
			req.User, timeTaken.Milliseconds(), err))
	}()

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	startTime := time.Now()

	user, err := s.uc.Get(ctx, req.Email)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - Get: %w", err))
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - Get: request { email: %s } completed in %dms with response { user: %#v, error: %#v }",
			req.Email, timeTaken.Milliseconds(), user, err))
	}()

	return &pb.GetUserResponse{
		User: &pb.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		},
	}, nil
}

func (s *GRPCServer) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.GetAllUsersResponse, error) {
	startTime := time.Now()

	users, err := s.uc.GetAll(ctx)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - GetAll: %w", err))
		return nil, status.Errorf(codes.Internal, "failed to get users: %v", err)
	}

	resp := make([]*pb.User, 0, len(users))
	for _, user := range users {
		resp = append(resp, &pb.User{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - GetAll: request { message: %#v } completed in %dms with response { users: %#v, error: %#v }",
			req, timeTaken.Milliseconds(), users, err))
	}()

	return &pb.GetAllUsersResponse{
		Users: resp,
	}, nil
}

func (s *GRPCServer) Update(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	startTime := time.Now()

	err := s.uc.Update(ctx, entity.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password,
		Role:     req.User.Role,
	})
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - Update: %w", err))
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - Update: request { user: %#v } completed in %dms with response { error: %#v }",
			req.User, timeTaken.Milliseconds(), err))
	}()

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	startTime := time.Now()

	err := s.uc.Delete(ctx, req.Email)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - Delete: %w", err))
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - Delete: request { email: %s } completed in %dms with response { error: %#v }",
			req.Email, timeTaken.Milliseconds(), err))
	}()

	return &emptypb.Empty{}, nil
}
