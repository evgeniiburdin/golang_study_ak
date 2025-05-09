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
		RPC    `yaml:"rpc"`
		Logger `yaml:"logger"`
		OpenCage
	}

	// App -.
	App struct {
		Name      string `yaml:"name"`
		Version   string `yaml:"version"`
		Transport string `yaml:"transport"`
	}

	// HTTP -.
	HTTP struct {
		Port       string `yaml:"port"`
		SwaggerURL string `yaml:"swaggerURL"`
	}

	// RPC -.
	RPC struct {
		Port string `yaml:"port"`
	}

	// Logger -.
	Logger struct {
		Level string `yaml:"level"`
	}

	// OpenCage -.
	OpenCage struct {
		APIKey string `env:"openCageAPIKey"`
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
