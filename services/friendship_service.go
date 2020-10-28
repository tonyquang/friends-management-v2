package services

import (
	"friends_management_v2/models/db_models"
	"friends_management_v2/models/request"
	"friends_management_v2/models/respone"
	"friends_management_v2/repository"
	"net/http"

	"gorm.io/gorm"
)

type Services interface {
	CreateNewUser(requestUser request.RequestCreateUser) *respone.ResponeError
}

// Manager is the implementation of recurring service
type Manager struct {
	dbconn *gorm.DB
}

// NewManager initializes recurring service
func NewManager(dbconn *gorm.DB) *Manager {
	return &Manager{
		dbconn: dbconn,
	}
}

func (m *Manager) CreateNewUser(requestUser request.RequestCreateUser) *respone.ResponeError {

	emailAddress := requestUser.Email
	password := requestUser.Password

	if ValidateEmail(emailAddress) == false {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusBadRequest, Description: "Email is Invalid!"}
	}

	if password == "" {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusBadRequest, Description: "Password is empty!"}
	}

	rs := repository.InsertNewUser(m.dbconn, db_models.Users{Email: emailAddress, Password: password})

	return rs
}
