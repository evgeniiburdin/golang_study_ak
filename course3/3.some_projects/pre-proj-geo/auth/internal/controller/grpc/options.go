package grpc

import (
	"net"
	"time"
)

// Option -.
type Option func(*GRPCServer)

// Port -.
func Port(port string) Option {
	return func(s *GRPCServer) {
		lis, err := net.Listen("tcp", ":"+port)
		if err != nil {
			s.lg.Error("grpc - listener - Port:", err)
			return
		}
		s.listener = lis
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *GRPCServer) {
		s.shutdownTimeout = timeout
	}
}
