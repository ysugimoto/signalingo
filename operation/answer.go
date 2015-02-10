package operation

import (
	"encoding/json"
)

type Answer struct {
	Type   string            `json:"type"`
	Sdp    string            `json:"sdp"`
	Sender string            `json:"sender"`
	Extra  map[string]string `json:"extra"`
}

func NewAnswerMessage(from, sdp string, extra map[string]string) ([]byte, error) {
	answer := Answer{
		Type:   CONNECTED,
		Sender: from,
		Sdp:    sdp,
		Extra:  extra,
	}

	return json.Marshal(answer)
}
