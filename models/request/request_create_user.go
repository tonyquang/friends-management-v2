package request

type RequestCreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
