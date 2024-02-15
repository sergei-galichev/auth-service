package env

import (
	"auth-service/internal/config"
	"github.com/pkg/errors"
	"net"
	"os"
)

const (
	cacheHostEnvName = "CACHE_HOST"
	cachePortEnvName = "CACHE_PORT"
)

type redisConfig struct {
	host string
	port string
}

func NewRedisConfig() (config.RedisConfig, error) {
	host := os.Getenv(cacheHostEnvName)
	port := os.Getenv(cachePortEnvName)

	if len(host) == 0 || host == "" {
		return nil, errors.New("redis hostname is empty or not found")
	}

	if len(port) == 0 || port == "" {
		return nil, errors.New("redis port is empty or not found")
	}

	return &redisConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
