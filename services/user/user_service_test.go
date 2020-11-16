package user

import (
	"testing"

	"friends_management_v2/utils"

	randomData "github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
)

// ================= BEGIN TEST func CreateNewUser ==============

func TestCreateNewUserSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	userMana := NewUserManager(tx)
	assert.NoError(t, userMana.CreateNewUser(Users{Email: randomData.Email()}))
}

func TestCreateNewUserExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	user := Users{Email: randomData.Email()}

	userMana := NewUserManager(tx)
	assert.NoError(t, userMana.CreateNewUser(user))
	assert.EqualError(t, userMana.CreateNewUser(user), "User is already exists!")
}

// ================= END TEST func CreateNewUser ==============

// ================= BEGIN TEST func GetListUser ==============

func TestGetListUserSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	userMana := NewUserManager(dbconn)

	actualRs, _ := userMana.GetListUser()
	assert.NotNil(t, actualRs)
}

// ================= END TEST func GetListUser ==============

// ================= BEGIN TEST func CheckUserExist ==============
func TestCheckUserExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	user := Users{Email: randomData.Email()}

	userMana := NewUserManager(tx)
	assert.NoError(t, userMana.CreateNewUser(user))

	actualRs, err := userMana.CheckUserExist([]string{user.Email})
	assert.Equal(t, true, actualRs)
	assert.Nil(t, err)
}

// ================= END TEST func CheckUserExist ==============
