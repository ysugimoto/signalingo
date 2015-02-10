package operation

import (
	"encoding/json"
)

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewErrorMessage(errorType, message string) ([]byte, error) {
	err := Error{
		Type:    errorType,
		Message: message,
	}

	return json.Marshal(err)
}
