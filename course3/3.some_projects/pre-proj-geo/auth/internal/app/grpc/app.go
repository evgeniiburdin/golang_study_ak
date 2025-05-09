// Package app_grpc configures and runs the application on a gRPC Server
package app_grpc

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"auth-service/config"
	grpccontroller "auth-service/internal/controller/grpc"
	"auth-service/internal/usecase"
	"auth-service/internal/usecase/repo"
	pb "auth-service/internal/usecase/user-service/user"
	jwtpkg "auth-service/pkg/jwt"
	"auth-service/pkg/logger"
	"auth-service/pkg/postgres"
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

	// User Service Client
	conn, err := grpc.Dial(cfg.App.UserServiceURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - couldn't connect to userService: %w", err))
	}
	defer conn.Close()
	usc := pb.NewUserServiceClient(conn)

	/*// Kafka Producer
	kafkaProducer, err := kafka_pkg.New(cfg.Kafka.Address, 20, time.Second*3, lg)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - couldn't create kafka producer: %w", err))
	}*/

	// Use case
	addressUseCase := usecase.New(usc, repo.New(pg), &jwtpkg.JWTService{})

	// GRPC Server
	grpcServer := grpccontroller.NewGRPCServer(addressUseCase, lg, grpccontroller.Port(cfg.RPC.Port))
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
