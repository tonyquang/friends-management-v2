package user

import (
	"errors"
	"testing"

	"friends_management_v2/utils"

	randomData "github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	const numsUser int = 1
	listUsers := []Users{}
	for i := 0; i < numsUser; i++ {
		listUsers = append(listUsers, Users{Email: randomData.Email()})
	}

	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()

	userMana := NewUserManager(tx)

	tcs := []struct {
		scenario      string
		mockInput     Users
		expectedError error
	}{
		{
			scenario:      "success",
			mockInput:     listUsers[0],
			expectedError: nil,
		},
		{
			scenario:      "User Exist",
			mockInput:     listUsers[0],
			expectedError: errors.New("User is already exists!"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.scenario, func(t *testing.T) {
			actualRs := userMana.CreateNewUser(tc.mockInput)
			assert.Equal(t, tc.expectedError, actualRs)
		})
	}
}

func TestGetListUserSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	userMana := NewUserManager(dbconn)

	actualRs, _ := userMana.GetListUser()
	assert.NotNil(t, actualRs)
}

func TestCheckUserExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()

	user := Users{Email: randomData.Email()}

	userMana := NewUserManager(tx)
	assert.NoError(t, userMana.CreateNewUser(user))

	actualRs, err := userMana.CheckUserExist([]string{user.Email})
	assert.Equal(t, true, actualRs)
	assert.Nil(t, err)
}
