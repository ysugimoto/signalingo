package operation

import (
	"encoding/json"
)

type Lock struct {
	Type string `json:"type"`
	User string `json:"user"`
}

func NewLockMessage(userId string) ([]byte, error) {
	lock := Lock{
		Type: CONNECTED,
		User: userId,
	}

	return json.Marshal(lock)
}
