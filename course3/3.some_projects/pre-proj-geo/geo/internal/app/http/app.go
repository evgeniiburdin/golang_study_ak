// Package app_http configures and runs the application on an HTTP Server
package app_http

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	"geo-service/config"
	v1 "geo-service/internal/controller/http/v1"
	"geo-service/internal/usecase"
	"geo-service/internal/usecase/webapi"
	"geo-service/pkg/httpserver"
	"geo-service/pkg/logger"
	"geo-service/pkg/metrics"
)

// Run -.
func Run(cfg *config.Config) {
	lg := logger.New(cfg.Logger.Level)

	// Use case
	addressUseCase := usecase.New(webapi.New(cfg.OpenCage.APIKey))

	// Metrics
	metricsService, err := metrics.NewMetricsService()
	if err != nil {
		log.Fatalf("app - Run - metrics: %v", err)
	}

	// HTTP Server
	httpRouter := chi.NewRouter()
	v1.NewRouter(httpRouter, cfg.HTTP.SwaggerURL, lg, addressUseCase, metricsService)
	httpServer := httpserver.New(httpRouter, httpserver.Port(cfg.HTTP.Port))
	lg.Info("http listening on " + cfg.HTTP.Port)

	// Waiting for signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lg.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		lg.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		lg.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
