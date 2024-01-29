package redis

import (
	"auth-service/pkg/cache"
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var _ cache.Repository = (*repository)(nil)

type repository struct {
	*sync.RWMutex
	client *redis.Client
}

func NewCacheRepository(size int) cache.Repository {

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-cache:6379",
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = ctx

	return &repository{
		client: client,
	}
}

func (r *repository) GetIterator() cache.Iterator {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Get(key []byte) ([]byte, error) {
	ctx := context.Background()

	r.RLock()
	defer r.RUnlock()

	return r.client.Get(ctx, string(key)).Bytes()
}

func (r *repository) Set(key []byte, value []byte, duration time.Duration) error {
	ctx := context.Background()

	r.Lock()
	defer r.Unlock()

	return r.client.Set(ctx, string(key), value, duration).Err()
}

func (r *repository) Del(key []byte) error {
	ctx := context.Background()

	r.Lock()
	defer r.Unlock()

	return r.client.Del(ctx, string(key)).Err()
}

func (r *repository) EntryCount() (entryCount int64) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) HitCount() (hitCount int64) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) MissCount() (missCount int64) {
	//TODO implement me
	panic("implement me")
}
