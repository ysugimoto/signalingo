package operation

import (
	"encoding/json"
)

type Candidate struct {
	Type      string `json:"type"`
	Candidate string `json:"candidate"`
	Sender    string `json:"sender"`
}

func NewCandidateMessage(from, candidate string) ([]byte, error) {
	c := Candidate{
		Type:      CONNECTED,
		Candidate: candidate,
		Sender:    from,
	}

	return json.Marshal(c)
}
