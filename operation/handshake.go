package operation

import (
	"encoding/json"
)

type Handshake struct {
	Type  string   `json:"type"`
	UUID  string   `json:"uuid"`
	Users []string `json:"users"`
}

func NewHandshakeMessage(userId string, users []string) ([]byte, error) {
	handshake := Handshake{
		Type:  CONNECTED,
		UUID:  userId,
		Users: users,
	}

	return json.Marshal(handshake)
}
