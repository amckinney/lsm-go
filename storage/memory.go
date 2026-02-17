package storage

import (
	"context"
	"sync"
)

// memoryStorage is a simple in-memory implementation of the Storage interface.
// It uses a map protected by a RWMutex for thread-safe access.
type memoryStorage struct {
	mu   sync.RWMutex
	data map[string][]byte
}

// NewMemoryStorage creates a new in-memory storage implementation.
func NewMemoryStorage() Storage {
	return &memoryStorage{
		data: make(map[string][]byte),
	}
}

// Get retrieves the value associated with the given key.
func (m *memoryStorage) Get(ctx context.Context, key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.data[key]
	if !ok {
		return nil, nil
	}

	// Return a copy to prevent external modification
	result := make([]byte, len(value))
	copy(result, value)
	return result, nil
}

// Set stores a value for the given key.
func (m *memoryStorage) Set(ctx context.Context, key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Store a copy to prevent external modification
	valueCopy := make([]byte, len(value))
	copy(valueCopy, value)
	m.data[key] = valueCopy
	return nil
}

// Close releases any resources held by the storage.
func (m *memoryStorage) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data = nil
	return nil
}
