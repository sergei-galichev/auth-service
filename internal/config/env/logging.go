package env

import (
	"auth-service/internal/config"
	"errors"
	"os"
)

const (
	loggingLevelEnvName = "LOG_LEVEL"
)

type loggingConfig struct {
	level string
}

func NewLoggingConfig() (config.LoggingConfig, error) {

	level := os.Getenv(loggingLevelEnvName)
	if len(level) == 0 {
		return nil, errors.New("logging level not found")
	}

	return &loggingConfig{
		level: level,
	}, nil
}

func (cfg *loggingConfig) LoggingLevel() string {
	return cfg.level
}
