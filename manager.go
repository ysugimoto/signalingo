package signaling

type Manager struct {
	connections map[string]*Connection
}

func NewManager() *Manager {
	return &Manager{
		connections: map[string]*Connection{},
	}
}

func (m *Manager) addConnection(conn *Connection) {
	if _, ok := m.connections[conn.UUID]; !ok {
		m.connections[conn.UUID] = conn
	}
}

func (m *Manager) removeConnection(conn *Connection) {
	if _, ok := m.connections[conn.UUID]; ok {
		delete(m.connections, conn.UUID)
	}
}
