package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Environment struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

type Config struct {
	Development Environment `yaml:"development"`
	Testing     Environment `yaml:"testing"`
	Production  Environment `yaml:"production"`
}

// Load reads config.yaml and returns the Environment configuration
// for the given env. env should be one of "development", "testing", or
// "production". Defaults to development if env is empty or unrecognized.
func Load(env string) (*Environment, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	switch env {
	case "testing":
		return &cfg.Testing, nil
	case "production":
		return &cfg.Production, nil
	default:
		return &cfg.Development, nil
	}
}
