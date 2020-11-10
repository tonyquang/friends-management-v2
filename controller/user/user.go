package user

type ResponeListUser struct {
	ListUsers []string `json:"list_users" binding:"required"`
	Count     uint64   `json:"count" binding:"required"`
}
type RequestCreateUser struct {
	Email string `json:"email" binding:"required"`
}
