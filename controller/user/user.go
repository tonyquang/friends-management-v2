package user

type ResponeListUser struct {
	ListUsers []string `json:"list_users" binding:"required"`
	Count     uint     `json:"count" binding:"required"`
}

type RequestCreateUser struct {
	Email string `json:"email" binding:"required"`
}

type HTTPSuccess struct {
	Success bool `json:"success" example:true`
}

type HTTPError struct {
	Message string `json:"error" example:"any error"`
}
