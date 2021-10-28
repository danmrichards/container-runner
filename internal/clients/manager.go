package clients

import (
	"strings"

	"github.com/google/uuid"
)

// Manager stores clients
type Manager struct {
	clients []string
}

// Register registers a new client.
func (m *Manager) Register() string {
	id := uuid.NewString()
	m.clients = append(m.clients, id)

	return id
}

func (m *Manager) String() string {
	if len(m.clients) == 0 {
		return "no clients"
	}
	return strings.Join(m.clients, ",")
}

// NewManager returns a new Manager.
func NewManager() *Manager {
	return &Manager{
		clients: make([]string, 0, 100),
	}
}
