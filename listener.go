package signaling

import (
	"golang.org/x/net/websocket"
)

func NewWebSocketListener() websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {

	})
}
