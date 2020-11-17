package common_respone

type HTTPSuccess struct {
	Success bool `json:"success" example:true`
}

type HTTPError struct {
	Message string `json:"error" example:"any error"`
}
