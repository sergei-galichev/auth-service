package config

import (
	"time"
)

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type LoggingConfig interface {
	LoggingLevel() string
}

type AuthConfig interface {
	Secret() string
	AdminKey() string
	AccessTTL() time.Duration
	RefreshTTL() time.Duration
}
