// Package app_jsonrpc configures and runs the application on a JSONRPC Server
package app_jsonrpc

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"geo-service/config"
	"geo-service/internal/controller/jsonrpc"
	"geo-service/internal/usecase"
	"geo-service/internal/usecase/webapi"
	"geo-service/pkg/logger"
)

// Run -.
func Run(cfg *config.Config) {
	lg := logger.New(cfg.Logger.Level)

	// Use case
	addressUseCase := usecase.New(webapi.New(cfg.OpenCage.APIKey))

	// JSONRPC Server
	_, err := jsonrpc.NewJSONRPCServer(addressUseCase, lg, jsonrpc.Port(cfg.RPC.Port))
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - jsonrpc.New: %w", err))
	}
	lg.Info("jsonrpc listening on " + cfg.RPC.Port)

	// Waiting for signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lg.Info("app - Run - signal: " + s.String())
	}
}
