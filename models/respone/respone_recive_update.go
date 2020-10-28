package respone

type ResponeReciveUpdate struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
