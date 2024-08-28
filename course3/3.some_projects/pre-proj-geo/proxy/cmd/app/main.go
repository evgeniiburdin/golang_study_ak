package main

import (
	"log"

	"geo-service-proxy/config"
	apphttp "geo-service-proxy/internal/app/http"
)

const (
	transportTypeHTTP = "http"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	switch cfg.App.Transport {
	case transportTypeHTTP:
		apphttp.Run(cfg)
	default:
		log.Fatalf("Unknown transport type: %s", cfg.App.Transport)
	}
}
