package main

import (
	"log"

	"geo-service/config"
	appgrpc "geo-service/internal/app/grpc"
	apphttp "geo-service/internal/app/http"
	appjsonrpc "geo-service/internal/app/jsonrpc"
	apprpc "geo-service/internal/app/rpc"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	switch cfg.App.Transport {
	case "rpc":
		apprpc.Run(cfg)
	case "jsonrpc":
		appjsonrpc.Run(cfg)
	case "grpc":
		appgrpc.Run(cfg)
	case "http":
		apphttp.Run(cfg)
	default:
		log.Fatalf("Unknown transport type: %s", cfg.App.Transport)
	}
}
