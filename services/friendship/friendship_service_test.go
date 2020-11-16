package friendship

import (
	"friends_management_v2/services/user"
	"friends_management_v2/utils"
	"testing"

	randomData "github.com/Pallinder/go-randomdata"
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
	assert.Nil(t, difference(expectedListUsers, actualListUsers))
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
	assert.Nil(t, difference(expectedMutualFriendsList, actualMutualFriendsList))
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

	assert.Nil(t, difference(expectedMutualFriendsList, actualMutualFriendsList))
}

// ==================================== END TEST GetMutualFriendsList FUNC =================================

// ==================================== BEGIN TEST Subscribe FUNC =================================

func TestSubscribeIfBothWasFriendSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)
	const numUsers int = 2
	users, ok := InsertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]}))

	actualRs := friendshipMana.Subscribe(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]})
	assert.Nil(t, actualRs)
}

func TestSubscribeIfBothWasNotFriendSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)
	const numUsers int = 2
	users, ok := InsertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	actualRs := friendshipMana.Subscribe(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]})
	assert.Nil(t, actualRs)
}

func TestSubscribeUserNotExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)
	actualRs := friendshipMana.Subscribe(FrienshipServiceInput{RequestEmail: "usernotexist@notfound.com", TargetEmail: "usernotexist2@notfound.com"})
	assert.EqualError(t, actualRs, "User Not Exist")
}

// ==================================== END TEST Subscribe FUNC =================================

// ==================================== BEGIN TEST Block FUNC =================================

func TestBlockIfBothWasFriendSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)
	const numUsers int = 2
	users, ok := InsertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]}))

	actualRs := friendshipMana.Block(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]})
	assert.Nil(t, actualRs)
}

func TestBlockIfBothWasNotFriendSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)
	const numUsers int = 2
	users, ok := InsertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	//assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]}))

	actualRs := friendshipMana.Block(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[1]})
	assert.Nil(t, actualRs)
}

func TestBlockUserNotExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)
	actualRs := friendshipMana.Block(FrienshipServiceInput{RequestEmail: "usernotexist@notfound.com", TargetEmail: "usernotexist2@notfound.com"})
	assert.EqualError(t, actualRs, "User Not Exist")
}

// ==================================== END TEST Block FUNC =================================

// ==================================== BEGIN TEST GetUsersReceiveUpdate FUNC =================================
func TestGetUsersReceiveUpdateSuccess(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)

	// Sender
	sender, ok0 := InsertUsersTest(tx, 1)
	assert.Equal(t, true, ok0)
	assert.Equal(t, 1, len(sender))

	// User will be use Make Friend with sender
	const numUsersMakeFriend int = 3
	usersWillMakeFriend, ok1 := InsertUsersTest(tx, numUsersMakeFriend)
	assert.Equal(t, true, ok1)
	assert.Equal(t, numUsersMakeFriend, len(usersWillMakeFriend))

	// User will be use subscribe to sender
	const numUsersSubscribe int = 3
	usersSubcribe, ok2 := InsertUsersTest(tx, numUsersSubscribe)
	assert.Equal(t, true, ok2)
	assert.Equal(t, numUsersSubscribe, len(usersSubcribe))

	// User mentioned
	const numUsersMentioned int = 2
	usersMentioned, ok3 := InsertUsersTest(tx, numUsersMentioned)
	assert.Equal(t, true, ok3)
	assert.Equal(t, numUsersMentioned, len(usersMentioned))

	// Make Friend
	for i := 0; i < numUsersMakeFriend; i++ {
		assert.NoError(t, friendshipMana.MakeFriend(FrienshipServiceInput{RequestEmail: usersWillMakeFriend[i], TargetEmail: sender[0]}))
	}

	// Subscribe
	for i := 0; i < numUsersSubscribe; i++ {
		assert.NoError(t, friendshipMana.Subscribe(FrienshipServiceInput{RequestEmail: usersSubcribe[i], TargetEmail: sender[0]}))
	}

	// Expected result
	expectedRs := []string{}
	expectedRs = append(expectedRs, usersWillMakeFriend...)
	expectedRs = append(expectedRs, usersSubcribe...)
	expectedRs = append(expectedRs, usersMentioned...)

	actualRs, err := friendshipMana.GetUsersReceiveUpdate(sender[0], usersMentioned)

	assert.NoError(t, err)
	assert.Nil(t, difference(actualRs, expectedRs))
}

func TestGetUsersReceiveUpdateUserNotExist(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipMana := NewFriendshipManager(tx)
	_, err := friendshipMana.GetUsersReceiveUpdate("usernotexist@notfound.com", []string{""})
	assert.EqualError(t, err, "User Not Exist")
}

// ==================================== END TEST GetUsersReceiveUpdate FUNC =================================

// InsertUsersTest
func InsertUsersTest(tx *gorm.DB, numsUser int) ([]string, bool) {
	listUsers := []string{}
	userMana := user.NewUserManager(tx)
	for i := 0; i < numsUser; i++ {
		email := randomData.Email()
		err := userMana.CreateNewUser(user.Users{Email: email})
		if err != nil {
			return nil, false
		}
		listUsers = append(listUsers, email)
	}
	return listUsers, true
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
