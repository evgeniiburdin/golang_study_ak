// Package app_grpc configures and runs the application on a gRPC Server
package app_grpc

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"user-service/config"
	"user-service/internal/controller/grpc"
	"user-service/internal/usecase"
	"user-service/internal/usecase/cache"
	"user-service/internal/usecase/repo"
	"user-service/pkg/logger"
	"user-service/pkg/postgres"
	redis_pkg "user-service/pkg/redis"
)

// Run -.
func Run(cfg *config.Config) {
	lg := logger.New(cfg.Logger.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Cache
	r, err := redis_pkg.New(cfg.Redis.Addr)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - redis_pkg.New: %w", err))
	}
	defer r.Close()
	cache, err := cache.New(r)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - cache.New: %w", err))
	}

	// Use case
	UserUseCase := usecase.New(repo.New(pg), cache)

	// GRPC Server
	grpcServer := grpc.NewGRPCServer(UserUseCase, lg, grpc.Port(cfg.RPC.Port))
	lg.Info("grpc listening on " + cfg.RPC.Port)

	/*// Kafka Server
	kafkaConsumer, err := kafka_pkg.New(cfg.Kafka.Address, "UserService")
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - kafka.New: %w", err))
	}
	kafkaServer := kafka_controller.New(UserUseCase, lg, *kafkaConsumer)
	err = kafkaServer.Start(15, time.Second*2)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - kafkaServer.Start: %w", err))
	}*/

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
	//kafkaConsumer.Close()
}
