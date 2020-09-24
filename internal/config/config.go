package config

import (
	"fmt"
	"os"
	"time"
)

// Config represent application configuration
type Config struct {
	Address     string
	StopTimeout time.Duration
	LogLevel    string
}

const (
	envKeyIP          = "GOPENCOV_IP"
	envKeyPort        = "GOPENCOV_PORT"
	envKeyStopTimeout = "GOPENCOV_STOP_TIMEOUT"
	envLogLevel       = "GOPENCOV_LOG_LEVEL"
)

// Collect collects application configuration
func Collect() (Config, error) {
	stopTimeout, err := time.ParseDuration(getenv(envKeyStopTimeout, "1m"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid stop timeout format: %w", err)
	}
	return Config{
		Address:     fmt.Sprintf("%s:%s", getenv(envKeyIP, "0.0.0.0"), getenv(envKeyPort, "4000")),
		StopTimeout: stopTimeout,
		LogLevel:    getenv(envLogLevel, "info"),
	}, nil
}

func getenv(key, defalt string) string {
	v := os.Getenv(key)
	if v == "" {
		return defalt
	}
	return v
}
