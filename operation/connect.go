package operation

import (
	"encoding/json"
)

type Connect struct {
	Type string `json:"type"`
	User string `json:"user"`
}

func NewConnectMessage(userId string) ([]byte, error) {
	connect := Connect{
		Type: CONNECTED,
		User: userId,
	}

	return json.Marshal(connect)
}
