package main

import (
	"log"

	"auth-service/config"
	app_grpc "auth-service/internal/app/grpc"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	switch cfg.App.Transport {
	case "grpc":
		app_grpc.Run(cfg)
	default:
		log.Fatalf("Unknown transport type: %s", cfg.App.Transport)
	}
}
