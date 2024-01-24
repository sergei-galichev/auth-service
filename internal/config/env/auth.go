package env

import (
	"auth-service/internal/config"
	"errors"
	"os"
	"time"
)

const (
	secretEnvName     = "SECRET_KEY"
	adminKeyEnvName   = "ADMIN_KEY"
	accessTTLEnvName  = "ACCESS_TTL"
	refreshTTLEnvName = "REFRESH_TTL"
)

type authConfig struct {
	secret     string
	adminKey   string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewAuthConfig() (config.AuthConfig, error) {
	secret := os.Getenv(secretEnvName)
	if len(secret) == 0 {
		return nil, errors.New("secret empty or not found")
	}

	adminKey := os.Getenv(adminKeyEnvName)
	if len(adminKey) == 0 {
		return nil, errors.New("admin empty or key not found")
	}

	aTTL := os.Getenv(accessTTLEnvName)
	if len(aTTL) == 0 {
		return nil, errors.New("access ttl not found")
	}

	accessTTL, err := time.ParseDuration(aTTL)
	if err != nil {
		return nil, errors.New("access ttl parse error")
	}

	rTTL := os.Getenv(refreshTTLEnvName)
	if len(rTTL) == 0 {
		return nil, errors.New("refresh ttl not found")
	}

	refreshTTL, err := time.ParseDuration(rTTL)
	if err != nil {
		return nil, errors.New("refresh ttl parse error")
	}

	return &authConfig{
		secret:     secret,
		adminKey:   adminKey,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}, nil
}

func (cfg *authConfig) Secret() string {
	return cfg.secret
}

func (cfg *authConfig) AdminKey() string {
	return cfg.adminKey
}

func (cfg *authConfig) AccessTTL() time.Duration {
	return cfg.accessTTL
}

func (cfg *authConfig) RefreshTTL() time.Duration {
	return cfg.refreshTTL
}
