package controllers

// Storage interface supports get and put of a single value
type Storage interface {
	Get() string
	Put(s string)
}

// MemStorage implements Storage
type MemStorage struct {
	value string
}

// Get our in-memory value
func (m *MemStorage) Get() string {
	return m.value
}

// Put our in-memory value
func (m *MemStorage) Put(s string) {
	m.value = s
}
