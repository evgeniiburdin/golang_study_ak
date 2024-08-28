package main

import (
	"fmt"
	"os"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getConfigFromYAML(data []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(data, &config)
	return config, err
}

func main() {
	file, err := os.Open("config.yaml")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file: %v", cerr)
		}
	}()
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}
	config, err := getConfigFromYAML(data)
	if err != nil {
		fmt.Printf("Error parsing config: %v", err)
		return
	}
	fmt.Println(config)
}
