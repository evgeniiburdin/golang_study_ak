package rpc

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"time"

	"geo-service/internal/entity"
	"geo-service/internal/usecase"
	"geo-service/pkg/logger"
)

const (
	defaultShutdownTimeout = 3 * time.Second
)

type RPCServer struct {
	listener        net.Listener
	notify          chan error
	shutdownTimeout time.Duration
	uc              usecase.Addresser
	lg              logger.Interface
}

func NewRPCServer(uc usecase.Addresser, lg logger.Interface, opts ...Option) (*RPCServer, error) {
	rpcServer := &RPCServer{
		listener:        nil,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
		uc:              uc,
		lg:              lg,
	}

	// Custom options
	for _, opt := range opts {
		opt(rpcServer)
	}

	err := rpc.Register(rpcServer)
	if err != nil {
		return nil, fmt.Errorf("rpc register error: %w", err)
	}

	go rpcServer.start()

	return rpcServer, nil
}

func (s *RPCServer) start() {
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.notify <- err
				return
			}
			go rpc.ServeConn(conn)
		}
	}()
}

// Notify -.
func (s *RPCServer) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *RPCServer) Shutdown() error {
	err := s.listener.Close()
	if err != nil {
		return fmt.Errorf("rpc shutdown error: %w", err)
	}
	return nil
}

func (s *RPCServer) GeocodeToAddress(args entity.Geocode, reply *entity.Address) error {
	startTime := time.Now()

	_, err := strconv.ParseFloat(args.Lat, 64)
	if err != nil {
		return err
	}
	_, err = strconv.ParseFloat(args.Lng, 64)
	if err != nil {
		return err
	}

	address, err := s.uc.GeocodeToAddress(context.Background(), entity.Geocode{
		Lat: args.Lat,
		Lng: args.Lng,
	})
	if err != nil {
		return fmt.Errorf("rpc - geocodeToAddess: %w", err)
	}

	reply.Country = address.Country
	reply.City = address.City

	defer func() {
		timeTaken := time.Since(startTime)
		s.lg.Info(fmt.Sprintf("grpc - geocodeToAddress: request { lat: %s lng: %s } completed in %dms with response { country: %s city: %s }",
			args.Lat, args.Lng, timeTaken.Milliseconds(), address.Country, address.City))
	}()

	return nil
}
