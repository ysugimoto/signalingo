package operation

type Users struct {
	UUID   string `json:"uuid"`
	Locked bool   `json:"locked"`
	Extra  string `json:"extra"`
}
