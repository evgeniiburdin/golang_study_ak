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
		RPC    `yaml:"rpc"`
		Logger `yaml:"logger"`
		PG     `yaml:"postgres"`
		Redis  `yaml:"redis"`
		Kafka
	}

	// App -.
	App struct {
		Name      string `yaml:"name"`
		Version   string `yaml:"version"`
		Transport string `yaml:"transport"`
	}

	// RPC -.
	RPC struct {
		Port string `yaml:"port"`
	}

	// Logger -.
	Logger struct {
		Level string `yaml:"level"`
	}

	// PG -.
	PG struct {
		URL     string `yaml:"url" env:"PG_URL"`
		PoolMax int    `yaml:"pool_max"`
	}

	// Redis -.
	Redis struct {
		Addr string `yaml:"address" env:"REDIS_ADDR"`
	}

	// Kafka -.
	Kafka struct {
		Address string `yaml:"address" env:"KAFKA_ADDRESS"`
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
