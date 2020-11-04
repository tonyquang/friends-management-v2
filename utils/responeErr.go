package utils

type ResponeErr struct {
	Error      error `json:"error"`
	StatusCode int   `json:"status_code"`
}
