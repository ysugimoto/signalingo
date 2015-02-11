package operation

import (
	"encoding/json"
)

type Offer struct {
	Type   string            `json:"type"`
	Sdp    string            `json:"sdp"`
	Sender string            `json:"sender"`
	Extra  map[string]string `json:"extra"`
}

func NewOfferMessage(from, sdp string, extra map[string]string) ([]byte, error) {
	offer := Offer{
		Type:   OFFER,
		Sender: from,
		Sdp:    sdp,
		Extra:  extra,
	}

	return json.Marshal(offer)
}
