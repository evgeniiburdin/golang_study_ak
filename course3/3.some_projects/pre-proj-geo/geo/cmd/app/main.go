package main

import (
	"log"

	"geo-service/config"
	appgrpc "geo-service/internal/app/grpc"
	apphttp "geo-service/internal/app/http"
	appjsonrpc "geo-service/internal/app/jsonrpc"
	apprpc "geo-service/internal/app/rpc"
)

const (
	transportTypeRPC     = "rpc"
	transportTypeJSONRPC = "jsonrpc"
	transportTypeGRPC    = "grpc"
	transportTypeHTTP    = "http"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	switch cfg.App.Transport {
	case transportTypeRPC:
		apprpc.Run(cfg)
	case transportTypeJSONRPC:
		appjsonrpc.Run(cfg)
	case transportTypeGRPC:
		appgrpc.Run(cfg)
	case transportTypeHTTP:
		apphttp.Run(cfg)
	default:
		log.Fatalf("Unknown transport type: %s", cfg.App.Transport)
	}
}
