package redis

import (
	"auth-service/pkg/cache"
	"auth-service/pkg/logging"
	"context"
	redisCache "github.com/go-redis/cache/v9"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var _ cache.Repository = (*repository)(nil)

type repository struct {
	sync.RWMutex
	rc *redisCache.Cache
}

func NewCache() cache.Repository {
	logger := logging.GetLogger()

	client := redis.NewClient(
		&redis.Options{
			Addr: "redis-cache:6379",
			DB:   0,
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Ping(ctx).Err()
	if err != nil {
		logger.Fatal("Failed to ping redis cache: ", err)
	}

	rc := redisCache.New(
		&redisCache.Options{
			Redis: client,
		},
	)

	return &repository{
		rc: rc,
	}
}

func (r *repository) Get(key []byte) ([]byte, error) {
	ctx := context.Background()
	var value []byte

	r.RLock()
	defer r.RUnlock()

	err := r.rc.Get(ctx, string(key), &value)
	if err != nil {
		return nil, errors.New("cache: error getting value")
	}

	return value, nil
}

func (r *repository) Set(key []byte, value []byte, duration time.Duration) error {
	ctx := context.Background()

	r.Lock()
	defer r.Unlock()
	item := &redisCache.Item{
		Ctx:   ctx,
		Key:   string(key),
		Value: value,
		TTL:   duration,
	}

	err := r.rc.Set(item)

	if err != nil {
		return errors.New("cache: error setting value")
	}

	return nil
}

func (r *repository) Del(key []byte) error {
	ctx := context.Background()

	r.Lock()
	defer r.Unlock()

	err := r.rc.Delete(ctx, string(key))
	if err != nil {
		return errors.New("cache: error deleting value")
	}

	return nil
}

func (r *repository) HitCount() (hitCount uint64) {
	return r.rc.Stats().Hits
}

func (r *repository) MissCount() (missCount uint64) {
	return r.rc.Stats().Misses
}
