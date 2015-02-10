package operation

// Operation signatures
const (
	HANDSHAKE    = "handshake"
	CONNECTED    = "connect"
	DISCONNECTED = "disconnect"
	LOCK         = "lock"
	UNLOCK       = "unlock"
	OFFER        = "offser"
	ANSWER       = "answer"
	CANDIDATE    = "candidate"

	LOCK_ERROR   = "lock_error"
	UNLOCK_ERROR = "unlock_error"
	OFFER_ERROR  = "offer_error"
	ANSWER_ERROR = "answer_error"

	USER_NOTFOUND = "user_not_found"
)

type Operation struct {
	Type      string            `json:"type"`
	Sdp       string            `json:"sdp"`
	Candidate string            `json:"candidate"`
	Sender    string            `json:"sender"`
	Target    string            `json:"target"`
	Extra     map[string]string `json:"extra"`
}
