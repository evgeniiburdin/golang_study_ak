package rpc

import (
	"net"
	"time"
)

// Option -.
type Option func(*RPCServer)

// Port -.
func Port(port string) Option {
	return func(s *RPCServer) {
		lis, err := net.Listen("tcp", ":"+port)
		if err != nil {
			s.lg.Error("rpc - listener - Port:", err)
			return
		}
		s.listener = lis
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *RPCServer) {
		s.shutdownTimeout = timeout
	}
}
