package operation

import (
	"encoding/json"
)

type Disconnect struct {
	Type string `json:"type"`
	User string `json:"user"`
}

func NewDisconnectMessage(userId string) ([]byte, error) {
	disconnect := Disconnect{
		Type: CONNECTED,
		User: userId,
	}

	return json.Marshal(disconnect)
}
