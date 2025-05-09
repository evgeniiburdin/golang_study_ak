package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Logger `yaml:"logger"`
	}

	// App -.
	App struct {
		Name               string `yaml:"name"`
		Version            string `yaml:"version"`
		Transport          string `yaml:"transport"`
		GeoServiceAddress  string `yaml:"geo_service_address" env:"GEO_SERVICE_ADDRESS"`
		AuthServiceAddress string `yaml:"auth_service_address" env:"AUTH_SERVICE_ADDRESS"`
	}

	// HTTP -.
	HTTP struct {
		Port       string `yaml:"port"`
		SwaggerURL string `yaml:"swaggerURL"`
	}

	// Logger -.
	Logger struct {
		Level string `yaml:"level"`
	}
)

// NewConfig returns app config
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	err = godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing env: %w", err)
	}

	return cfg, nil
}
