package signaling

import (
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

type Connection struct {
	UUID  string
	Conn  *websocket.Conn
	State uint
}

func NewConnection(ws *websocket.Conn) *Connection {
	return &Connection{
		UUID:  uuid.NewV4().String(),
		Conn:  ws,
		State: 0,
	}
}
