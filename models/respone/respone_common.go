package respone

//using for common respone success
type ResponeSuccess struct {
	Success bool `json:"success"`
}

//using for all error respone
type ResponeError struct {
	StatusCode  int    `json:"code"`
	Success     bool   `json:"success"`
	Description string `json:"description"`
}
