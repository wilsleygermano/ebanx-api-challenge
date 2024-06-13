package event

type EventRequestBody struct {
	Type        string  `json:"type"`
	Destination string  `json:"destination"`
	Amount      float64 `json:"amount"`
	Origin      string  `json:"origin"`
}
