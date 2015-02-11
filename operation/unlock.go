package operation

import (
	"encoding/json"
)

type Unlock struct {
	Type string `json:"type"`
	User string `json:"user"`
}

func NewUnlockMessage(userId string) ([]byte, error) {
	unlock := Unlock{
		Type: UNLOCK,
		User: userId,
	}

	return json.Marshal(unlock)
}
