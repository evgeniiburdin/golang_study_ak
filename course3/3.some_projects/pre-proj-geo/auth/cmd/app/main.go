package main

import (
	"log"

	"auth-service/config"
	appgrpc "auth-service/internal/app/grpc"
)

const (
	transportTypeGRPC = "grpc"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	switch cfg.App.Transport {
	case transportTypeGRPC:
		appgrpc.Run(cfg)
	default:
		log.Fatalf("Unknown transport type: %s", cfg.App.Transport)
	}
}
