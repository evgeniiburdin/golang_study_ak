// Package app_rpc configures and runs the application on an RPC Server
package app_rpc

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"geo-service/config"
	rpcinternal "geo-service/internal/controller/rpc"
	"geo-service/internal/usecase"
	"geo-service/internal/usecase/webapi"
	"geo-service/pkg/logger"
)

// Run -.
func Run(cfg *config.Config) {
	lg := logger.New(cfg.Logger.Level)

	// Use case
	addressUseCase := usecase.New(webapi.New(cfg.OpenCage.APIKey))

	// RPC Server
	rpcServer, err := rpcinternal.NewRPCServer(addressUseCase, lg, rpcinternal.Port(cfg.RPC.Port))
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - rpcServer.New: %w", err))
	}
	lg.Info("rpc listening on " + cfg.RPC.Port)

	// Waiting for signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lg.Info("app - Run - signal: " + s.String())
	case err := <-rpcServer.Notify():
		lg.Error(fmt.Errorf("app - Run - rpcServer.Notify: %w", err))
	}

	err = rpcServer.Shutdown()
	if err != nil {
		lg.Error(fmt.Errorf("app - Run - rpcServer.Shutdown: %w", err))
	}
}
