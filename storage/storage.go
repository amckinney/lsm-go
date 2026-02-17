package storage

import "context"

// Storage is the core abstraction for the LSM storage engine.
// It provides a simple key-value interface for reading and writing data.
type Storage interface {
	// Get retrieves the value associated with the given key.
	// Returns nil if the key does not exist.
	Get(ctx context.Context, key string) ([]byte, error)

	// Set stores a value for the given key.
	// If the key already exists, it will be overwritten.
	Set(ctx context.Context, key string, value []byte) error

	// Close releases any resources held by the storage.
	Close() error
}
