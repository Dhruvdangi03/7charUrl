package store

import "sync"

type MemoryStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (m *MemoryStore) Save(short, original string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[short] = original
}

func (m *MemoryStore) Get(short string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[short]
	return val, ok
}
