// Package app_http configures and runs the application on an HTTP Server
package app_http

import (
	"fmt"
	pbauth "geo-service-proxy/api/auth-service/auth"
	pbgeo "geo-service-proxy/api/geo-service/geo"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"

	"geo-service-proxy/config"
	v1 "geo-service-proxy/internal/controller/http/v1"
	"geo-service-proxy/internal/usecase"
	"geo-service-proxy/pkg/httpserver"
	"geo-service-proxy/pkg/logger"
	"geo-service-proxy/pkg/metrics"
)

// Run -.
func Run(cfg *config.Config) {
	lg := logger.New(cfg.Logger.Level)

	// Geo-service grpc client
	geoConn, err := grpc.Dial(cfg.App.GeoServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("app run couldn't connect to geo-service via grpc: %v", err)
	}
	defer geoConn.Close()

	geoClient := pbgeo.NewGeoServiceClient(geoConn)

	// Auth-service grpc client
	authConn, err := grpc.Dial(cfg.App.AuthServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("app run couldn't connect to auth-service via grpc: %v", err)
	}
	defer authConn.Close()

	authClient := pbauth.NewAuthServiceClient(authConn)

	// Use case
	proxyUseCase := usecase.New(authClient, geoClient)

	// Metrics
	metricsService, err := metrics.NewMetricsService()
	if err != nil {
		log.Fatalf("app - Run - metrics: %v", err)
	}

	// HTTP Server
	httpRouter := chi.NewRouter()
	v1.NewRouter(httpRouter, cfg.HTTP.SwaggerURL, lg, proxyUseCase, metricsService)
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
