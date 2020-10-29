package services

import (
	"friends_management_v2/models/request"
	"friends_management_v2/models/respone"
	"friends_management_v2/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	db := utils.CreateConnection()
	testCases := []struct {
		name           string
		givenUser      request.RequestCreateUser
		expectedResult *respone.ResponeError // Expected when recive nil value, and error when recvie not nil value
	}{
		{
			name: "Success Created New User",
			givenUser: request.RequestCreateUser{
				Email:    "tonyquang123@gmail.com",
				Password: "Password",
			},
			expectedResult: nil,
		},
		{
			name: "Invalid Email",
			givenUser: request.RequestCreateUser{
				Email:    "a",
				Password: "b",
			},
			expectedResult: &respone.ResponeError{
				Success:     false,
				StatusCode:  http.StatusBadRequest,
				Description: "Email is Invalid!",
			},
		},
		{
			name: "Empty Password",
			givenUser: request.RequestCreateUser{
				Email:    "tonyquang@gmail.com",
				Password: "",
			},
			expectedResult: &respone.ResponeError{
				Success:     false,
				StatusCode:  http.StatusBadRequest,
				Description: "Password is empty!",
			},
		},
		{
			name: "Is Already User",
			givenUser: request.RequestCreateUser{
				Email:    "tonyquang123@gmail.com",
				Password: "pass1",
			},
			expectedResult: &respone.ResponeError{
				Success:     false,
				StatusCode:  http.StatusBadRequest,
				Description: "User is already!",
			},
		},
	}

	tx := db.Begin()

	assert.NoError(t, utils.LoadFixture(tx, "./datatest/create_user.sql"))

	tx.SavePoint("sp2")
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewManager(tx)

			rs := manager.CreateNewUser(tt.givenUser)

			//Then
			assert.Equal(t, tt.expectedResult, rs)
		})
	}
	tx.RollbackTo("sp2")
}
