package user

type ResponeListUser struct {
	ListUsers []string `json:"list_users"`
	Count     uint64   `json:"count"`
}
type RequestCreateUser struct {
	Email string `json:"email"`
}
