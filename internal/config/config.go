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
	DBDriver    string
	DBURI       string
}

const (
	envKeyIP          = "GOPENCOV_IP"
	envKeyPort        = "GOPENCOV_PORT"
	envKeyStopTimeout = "GOPENCOV_STOP_TIMEOUT"
	envLogLevel       = "GOPENCOV_LOG_LEVEL"
	envDBDriver       = "GOPENCOV_DB_DRIVER"
	envDBURI          = "GOPENCOV_DB_URI"
)

// Collect collects application configuration
func Collect() (Config, error) {
	stopTimeout, err := time.ParseDuration(getenv(envKeyStopTimeout, "1m"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid stop timeout format: %w", err)
	}

	dbDriver := getenv(envDBDriver, "")
	if dbDriver == "" {
		return Config{}, fmt.Errorf("enviroment variable %q is not set", envDBDriver)
	}

	dbURI := getenv(envDBURI, "")
	if dbURI == "" {
		return Config{}, fmt.Errorf("enviroment variable %q is not set", envDBURI)
	}

	return Config{
		Address:     fmt.Sprintf("%s:%s", getenv(envKeyIP, "0.0.0.0"), getenv(envKeyPort, "4000")),
		StopTimeout: stopTimeout,
		LogLevel:    getenv(envLogLevel, "info"),
		DBDriver:    dbDriver,
		DBURI:       dbURI,
	}, nil
}

func getenv(key, defalt string) string {
	v := os.Getenv(key)
	if v == "" {
		return defalt
	}
	return v
}
