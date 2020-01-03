package atomic

import "sync"

import "errors"

// SafeMap uses a mutex to allow
// getting and setting in a thread-safe way
type SafeMap struct {
	m  map[string]string
	mu *sync.RWMutex
}

// NewSafeMap creates a new SafeMap
func NewSafeMap() SafeMap {
	return SafeMap{
		m:  make(map[string]string),
		mu: &sync.RWMutex{},
	}
}

// Set uses a write lock and sets the value given a key
func (sm *SafeMap) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

// Get uses a Read lock and gets the value if it exists,
// otherwise an error is returned
func (sm *SafeMap) Get(key string) (string, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if value, ok := sm.m[key]; ok {
		return value, nil
	}
	return "", errors.New("Key not found")
}
