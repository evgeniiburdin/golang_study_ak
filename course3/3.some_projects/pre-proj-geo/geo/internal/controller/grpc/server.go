package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "geo-service/internal/controller/grpc/gen/geo-service/geo"
	"geo-service/internal/entity"
	"geo-service/internal/usecase"
	"geo-service/pkg/logger"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type GRPCServer struct {
	pb.UnimplementedGeoServiceServer
	server          *grpc.Server
	listener        net.Listener
	notify          chan error
	shutdownTimeout time.Duration
	uc              usecase.Addresser
	lg              logger.Interface
}

func NewGRPCServer(uc usecase.Addresser, lg logger.Interface, opts ...Option) *GRPCServer {
	grpcServer := &GRPCServer{
		server:          grpc.NewServer(),
		listener:        nil,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		uc:              uc,
		lg:              lg,
	}
	pb.RegisterGeoServiceServer(grpcServer.server, grpcServer)

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

func (s *GRPCServer) GeocodeToAddress(ctx context.Context, in *pb.Geocode) (*pb.Address, error) {
	startTime := time.Now()

	address, err := s.uc.GeocodeToAddress(
		ctx,
		entity.Geocode{
			Lat: in.Lat,
			Lng: in.Lng,
		},
	)
	if err != nil {
		s.lg.Error(err, fmt.Errorf("grpc - geocodeToAddress: %w", err))
		return nil, fmt.Errorf("grpc - geocodeToAddress: %w", err)
	}

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - geocodeToAddress: request { lat: %s lng: %s } completed in %dms with response { country: %s city: %s }",
			in.Lat, in.Lng, timeTaken.Milliseconds(), address.Country, address.City))
	}()

	return &pb.Address{
		Country: address.Country,
		City:    address.City,
	}, nil
}
