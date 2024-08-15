package main

import (
	"log"

	"geo-service/config"
	app_grpc "geo-service/internal/app/grpc"
	app_http "geo-service/internal/app/http"
	app_jsonrpc "geo-service/internal/app/jsonrpc"
	app_rpc "geo-service/internal/app/rpc"
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
		app_rpc.Run(cfg)
	case "jsonrpc":
		app_jsonrpc.Run(cfg)
	case "grpc":
		app_grpc.Run(cfg)
	case "http":
		app_http.Run(cfg)
	default:
		log.Fatalf("Unknown transport type: %s", cfg.App.Transport)
	}
}
