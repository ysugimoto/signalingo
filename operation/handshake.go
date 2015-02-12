package operation

import (
	"encoding/json"
)

type Handshake struct {
	Type  string            `json:"type"`
	UUID  string            `json:"uuid"`
	Users []Users           `json:"users"`
	Extra map[string]string `json:"extra"`
}

func NewHandshakeMessage(userId string, extra map[string]string, users []Users) ([]byte, error) {
	handshake := Handshake{
		Type:  HANDSHAKE,
		UUID:  userId,
		Users: users,
		Extra: extra,
	}

	return json.Marshal(handshake)
}
