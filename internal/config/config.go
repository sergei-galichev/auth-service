package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func InitDefaultConfig() {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	user := "postgres"
	password := "postgres"
	sslMode := "disable"
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host,
		port,
		dbName,
		user,
		password,
		sslMode,
	)
	viper.SetDefault("PG_DSN", dsn)
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("GRPC_HOST", "localhost")
	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("SECRET_KEY", "secret")
	viper.SetDefault("ADMIN_KEY", "admin")
	viper.SetDefault("ACCESS_TTL", time.Minute*5)
	viper.SetDefault("REFRESH_TTL", time.Minute*10)
}
