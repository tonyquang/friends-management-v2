package request

type RequestReciveUpdate struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
