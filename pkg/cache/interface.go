package cache

import "time"

type Repository interface {
	// GetIterator returns iterator for storage
	GetIterator() Iterator

	// Get returns value from storage or error if value not found
	Get(key []byte) ([]byte, error)

	// Set sets a key, value and ttl for storage entry and stores it.
	// expireIn <= 0 means no expires, but it can be evicted when storage is full
	Set(key []byte, value []byte, duration time.Duration) error

	// Del deletes an item in the storage by key and returns true or false if a delete operation occurred
	Del(key []byte) error

	// EntryCount returns number of items currently in the storage
	EntryCount() (entryCount int64)

	// HitCount is a metric that returns the number of times a key was found in the storage
	HitCount() (hitCount int64)

	// MissCount is a metric that returns the number of times a miss occurred in the storage
	MissCount() (missCount int64)
}
