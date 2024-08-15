package jsonrpc

import (
	"net"
	"time"
)

// Option -.
type Option func(*JSONRPCServer)

// Port -.
func Port(port string) Option {
	return func(s *JSONRPCServer) {
		lis, err := net.Listen("tcp", ":"+port)
		if err != nil {
			s.lg.Error("jsonrpc - listener - Port:", err)
			return
		}
		s.listener = lis
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *JSONRPCServer) {
		s.shutdownTimeout = timeout
	}
}
