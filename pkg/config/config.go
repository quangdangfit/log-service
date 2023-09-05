package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	ProductionEnv = "production"

	DatabaseTimeout  = 5 * time.Second
	DefaultPageLimit = 100
)

var (
	AuthIgnoreMethods []string
)

type Schema struct {
	Environment  string `env:"environment"`
	Port         int    `env:"port"`
	SecretAPIKey string `env:"secret_api_key"`
	DatabaseURI  string `env:"database_uri"`
	DatabaseName string `env:"database_name"`
}

var (
	cfg Schema
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	environment := os.Getenv("environment")
	err := godotenv.Load(filepath.Join(currentDir, "config.yaml"))
	if err != nil && environment == "localhost" {
		log.Fatalf("Error on load configuration file, error: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error on parsing configuration file, error: %v", err)
	}
}

func GetConfig() *Schema {
	return &cfg
}
