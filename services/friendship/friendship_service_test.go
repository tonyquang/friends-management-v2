package friendship

import (
	"fmt"
	"friends_management_v2/services/user"
	"friends_management_v2/utils"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// ==================================== BEGIN TEST MakeFriend FUNC =================================
// TestMakeFriendSuccess func test create friend connection between 2 users and both not yet subscribe/block together
func TestMakeFriendSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	// Insert Testing User to Database
	users, ok := InsertUsersTest(tx, 2)
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, len(users))

	friendshipMana := NewFriendshipManager(tx)
	assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]}))
}

func TestExistFriendship(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	// Insert Testing User to Database
	users, ok := InsertUsersTest(tx, 2)
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, len(users))

	friendshipMana := NewFriendshipManager(tx)
	assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]}))

	assert.EqualError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]}), "Friendship was exist")
}

func TestMakeFriendWithUserNotExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	// Insert Testing User to Database
	users, ok := InsertUsersTest(tx, 1)
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, len(users))

	usersNotExist := "usernotexist123@notexist.notfound"

	friendshipMana := NewFriendshipManager(tx)
	assert.EqualError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: usersNotExist}), "User Not Exist")
}

// ==================================== END TEST MakeFriend FUNC =================================

// ==================================== BEGIN TEST GetUserFriendList FUNC =================================

func TestGetUserFriendListSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	const numUsers int = 10
	users, ok := InsertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	friendshipMana := NewFriendshipManager(tx)
	for i := 1; i < numUsers; i++ {
		assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[i]}))
	}

	expectedListUsers := []string{}
	expectedListUsers = append(expectedListUsers, users[1:]...)
	actualListUsers, err := friendshipMana.GetUserFriendList(user.Users{Email: users[0]})

	assert.NoError(t, err)
	assert.NotNil(t, actualListUsers)
	assert.Equal(t, true, reflect.DeepEqual(expectedListUsers, actualListUsers))
}

func TestGetUserFriendListWithUserNotExist(t *testing.T) {
	dbconn := utils.CreateConnection()

	friendshipMana := NewFriendshipManager(dbconn)
	actualFriendsList, err := friendshipMana.GetUserFriendList(user.Users{Email: "usernotexist@notfound.com"})
	assert.Nil(t, actualFriendsList)
	assert.EqualError(t, err, "User Not Exist")
}

// ==================================== END TEST GetUserFriendList FUNC =================================

// ==================================== BEGIN TEST GetMutualFriendsList FUNC =================================

func TestGetMutualFriendsListSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	const numUsers int = 6
	users, ok := InsertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	first_user := users[0]
	second_user := users[1]

	friendshipMana := NewFriendshipManager(tx)
	for i := 2; i < numUsers; i++ {
		assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: first_user, TargetEmail: users[i]}))
		assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: second_user, TargetEmail: users[i]}))
	}

	expectedMutualFriendsList := []string{}
	expectedMutualFriendsList = append(expectedMutualFriendsList, users[2:]...)
	actualMutualFriendsList, err := friendshipMana.GetMutualFriendsList(FrienshipServiceInput{RequestEmail: first_user, TargetEmail: second_user})
	assert.NotNil(t, actualMutualFriendsList)
	assert.NoError(t, err)
	assert.Equal(t, true, reflect.DeepEqual(expectedMutualFriendsList, actualMutualFriendsList))
}

func TestGetMutualFriendsListUserNotExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	const numUsers int = 6
	users, ok := InsertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	first_user := users[0]
	second_user := users[1]

	friendshipMana := NewFriendshipManager(tx)
	for i := 2; i < numUsers; i++ {
		assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: first_user, TargetEmail: users[i]}))
		assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: second_user, TargetEmail: users[i]}))
	}

	expectedMutualFriendsList := []string{}
	expectedMutualFriendsList = append(expectedMutualFriendsList, users[2:]...)
	actualMutualFriendsList, err := friendshipMana.GetMutualFriendsList(FrienshipServiceInput{RequestEmail: first_user, TargetEmail: second_user})
	assert.NotNil(t, actualMutualFriendsList)
	assert.NoError(t, err)
	assert.Equal(t, true, reflect.DeepEqual(expectedMutualFriendsList, actualMutualFriendsList))
}

// ==================================== END TEST GetMutualFriendsList FUNC =================================

// InsertUsersTest
func InsertUsersTest(tx *gorm.DB, numsUser int) ([]string, bool) {
	listUsers := []string{}
	userMana := user.NewUserManager(tx)
	for i := 0; i < numsUser; i++ {
		email := "usertest" + fmt.Sprint(i) + "@usertest.com"
		err := userMana.CreateNewUser(user.Users{Email: email})
		if err != nil {
			return nil, false
		}
		listUsers = append(listUsers, email)
	}
	return listUsers, true
}
