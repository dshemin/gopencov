package logger

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strings"
)

// New make new instance of the logger
func New(level string) (*log.Logger, error) {
	lvl, err := stringToLvl(level)
	if err != nil {
		return nil, fmt.Errorf("cannot parse log level: %w", err)
	}
	logger := log.New(os.Args[0])
	logger.DisableColor()
	logger.SetLevel(lvl)
	return logger, nil
}

func stringToLvl(s string) (log.Lvl, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return log.DEBUG, nil
	case "info":
		return log.INFO, nil
	case "warn":
		return log.WARN, nil
	case "error":
		return log.ERROR, nil
	case "off":
		return log.OFF, nil
	}
	return log.Lvl(0), fmt.Errorf("unknown log level %q", s)
}
