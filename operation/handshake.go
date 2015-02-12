package operation

import (
	"encoding/json"
)

type Handshake struct {
	Type  string  `json:"type"`
	UUID  string  `json:"uuid"`
	Users []Users `json:"users"`
}

func NewHandshakeMessage(userId string, users []Users) ([]byte, error) {
	handshake := Handshake{
		Type: HANDSHAKE,
		UUID: userId,
	}

	return json.Marshal(handshake)
}
