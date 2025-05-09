// Package app_grpc configures and runs the application on a gRPC Server
package app_grpc

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"geo-service/config"
	"geo-service/internal/controller/grpc"
	"geo-service/internal/usecase"
	"geo-service/internal/usecase/webapi"
	"geo-service/pkg/logger"
)

// Run -.
func Run(cfg *config.Config) {
	lg := logger.New(cfg.Logger.Level)

	// Use case
	addressUseCase := usecase.New(webapi.New(cfg.OpenCage.APIKey))

	// GRPC Server
	grpcServer := grpc.NewGRPCServer(addressUseCase, lg, grpc.Port(cfg.RPC.Port))
	lg.Info("grpc listening on " + cfg.RPC.Port)

	// Waiting for signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lg.Info("app - Run - signal: " + s.String())
	case err := <-grpcServer.Notify():
		lg.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}

	grpcServer.Shutdown()
}
