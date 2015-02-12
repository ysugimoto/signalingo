package signaling

import (
	"errors"
	"github.com/satori/go.uuid"
	"github.com/ysugimoto/signalingo/operation"
	"golang.org/x/net/websocket"
	"log"
)

type Connection struct {
	UUID   string
	Conn   *websocket.Conn
	Closed bool
	Locked bool
	Extra  map[string]string
}

var manager = NewManager()

func NewWebSocketConnectionHandler(env Env) websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		client := NewConnection(ws)

		// Handshake message
		// Wait a few seconds
		if msg, err := operation.NewHandshakeMessage(client.UUID, client.Extra, manager.GetAllUsers()); err == nil {

			// Does need hook?
			if env.Hook.Url == "" {
				if err := client.Send(msg); err != nil {
					log.Println("Handshake send failed")
				} else {
					log.Printf("UUID: %s handshake", client.UUID)
				}
			} else {
				hook := NewHook(env.Hook.Url, client.Extra)
				hook.Run()
				resp := <-hook.Resp
				if resp == 0 {
					if err := client.Send(msg); err != nil {
						log.Println("Handshake send failed")
					} else {
						log.Printf("UUID: %s handshake", client.UUID)
					}
				} else {
					client.Close()
					return
				}
			}
		}

		manager.AddConnection(client)

		defer func() {
			if err := client.Close(); err == nil {
				manager.RemoveConnection(client)
			}
		}()

		for {
			if client.Closed {
				break
			}
			if msg, err := client.Receive(); err != nil {
				manager.RemoveConnection(client)
				break
			} else {
				log.Printf("[WebSocket] message: %s\n", msg)
				manager.HandleMessage(msg)
			}
		}

	})
}

func NewConnection(ws *websocket.Conn) *Connection {
	query := ws.Request().URL.Query()
	extra := make(map[string]string)
	for key, value := range query {
		extra[key] = value[0]
	}

	return &Connection{
		UUID:   uuid.NewV4().String(),
		Conn:   ws,
		Closed: false,
		Locked: false,
		Extra:  extra,
	}

}

func (c *Connection) Send(message []byte) (err error) {
	if c.Closed {
		return errors.New("Connection has already closed.")
	}

	return websocket.Message.Send(c.Conn, string(message))
}

func (c *Connection) Receive() (msg string, err error) {
	err = websocket.Message.Receive(c.Conn, &msg)
	return msg, err
}

func (c *Connection) Close() error {
	if err := c.Conn.Close(); err != nil {
		return err
	}

	c.Closed = true
	return nil
}
