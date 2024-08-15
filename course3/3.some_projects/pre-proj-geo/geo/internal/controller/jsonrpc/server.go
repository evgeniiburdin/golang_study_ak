package jsonrpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"time"

	"geo-service/internal/entity"
	"geo-service/internal/usecase"
	"geo-service/pkg/logger"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type JSONRPCServer struct {
	server          http.Server
	listener        net.Listener
	notify          chan error
	shutdownTimeout time.Duration
	uc              usecase.Addresser
	lg              logger.Interface
}

func NewJSONRPCServer(uc usecase.Addresser, lg logger.Interface, opts ...Option) (*JSONRPCServer, error) {
	jsonRPCServer := &JSONRPCServer{
		server:          http.Server{},
		listener:        nil,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		uc:              uc,
		lg:              lg,
	}

	// Custom options
	for _, opt := range opts {
		opt(jsonRPCServer)
	}

	err := rpc.Register(jsonRPCServer)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc register error: %w", err)
	}

	jsonRPCServer.start()

	return jsonRPCServer, nil
}

func (s *JSONRPCServer) start() {
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.lg.Error(fmt.Sprintf("jsonrpc connection accept error: %v", err))
				continue
			}

			go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}()
}

func (s *JSONRPCServer) GeocodeToAddress(args *entity.Geocode, reply *entity.Address) error {
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
		return fmt.Errorf("jsonrpc - geocodeToAddess: %w", err)
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
