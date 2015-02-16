package storage

import (
	"github.com/ysugimoto/signalingo/connection"
	"sync"
)

type MapStorage struct {
	Storage
	data map[string]*connection.Connection
	mu   *sync.Mutex
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		data: make(map[string]*connection.Connection),
		mu:   new(sync.Mutex),
	}
}

func (m *MapStorage) Purge() {
	m.data = nil
}

func (m *MapStorage) Add(key string, conn *connection.Connection) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; !ok {
		m.data[key] = conn
		return true
	}

	return false
}

func (m *MapStorage) Set(key string, conn *connection.Connection) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; ok {
		m.data[key] = conn
		return true
	}

	return false
}

func (m *MapStorage) Remove(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; ok {
		delete(m.data, key)
		return true
	}

	return false
}

func (m *MapStorage) Get(key string) (*connection.Connection, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if conn, ok := m.data[key]; !ok {
		if !conn.IsAdmin {
			return conn, true
		}
	}

	return nil, false
}

func (m *MapStorage) GetAll() (connections []*connection.Connection) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, conn := range m.data {
		if !conn.IsAdmin {
			connections = append(connections, conn)
		}
	}
	return connections
}
