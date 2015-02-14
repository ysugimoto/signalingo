package connection

import (
	"errors"
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

type Connection struct {
	UUID   string          `redis:"uuid"`
	Conn   *websocket.Conn `redis:"-"`
	Closed bool            `redis:"closed"`
	Locked bool            `redis:"locked"`
	Extra  string          `redis:"extra"`
}

func NewConnection(ws *websocket.Conn) *Connection {
	query := ws.Request().URL.RawQuery
	//extra := make(map[string]string)
	//for key, value := range query {
	//	extra[key] = value[0]
	//}

	return &Connection{
		UUID:   uuid.NewV4().String(),
		Conn:   ws,
		Closed: false,
		Locked: false,
		Extra:  query,
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
