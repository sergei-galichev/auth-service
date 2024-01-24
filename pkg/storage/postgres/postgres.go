package postgres

import (
	"auth-service/pkg/logging"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"time"
)

type Storage struct {
	Pool *pgxpool.Pool
}

func New(dsn string) (*Storage, error) {
	logger := logging.GetLogger()
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Error("Failed to parse postgres config: ", err)
		return nil, errors.WithStack(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Error("Failed to create postgres pool: ", err)
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = pool.Ping(ctx); err != nil {
		logger.Error("Failed to ping postgres: ", err)
	}

	logger.Info("Postgres is connected")

	return &Storage{
		Pool: pool,
	}, nil
}

func (s *Storage) Close() {
	s.Pool.Close()
}
