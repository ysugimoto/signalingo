package operation

import (
	"encoding/json"
)

type Shutdown struct {
	Type string `json:"type"`
}

func NewShutdownMessage() ([]byte, error) {
	shutdown := Shutdown{
		Type: SHUTDOWN,
	}

	return json.Marshal(shutdown)
}
