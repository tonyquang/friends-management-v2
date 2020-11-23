package friendship

import (
	"errors"
	"testing"

	"friends_management_v2/services/user"
	"friends_management_v2/utils"

	randomData "github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestMakeFriendSuccess func test create friend connection between 2 users and both not yet subscribe/block together
func TestMakeFriend(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	users, ok := insertUsersTest(tx, 2)
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, len(users))
	friendshipManager := NewFriendshipManager(tx)
	testCase := []struct {
		scenario      string
		mockInput     FrienshipServiceInput
		expectedError error
	}{
		{
			scenario: "Success",
			mockInput: FrienshipServiceInput{
				RequestEmail: users[0],
				TargetEmail:  users[1],
			},
			expectedError: nil,
		},
		{
			scenario: "Exist Friendship",
			mockInput: FrienshipServiceInput{
				RequestEmail: users[0],
				TargetEmail:  users[1],
			},
			expectedError: errors.New("Friendship was exist"),
		},
		{
			scenario: "User not exist",
			mockInput: FrienshipServiceInput{
				RequestEmail: users[0],
				TargetEmail:  "usernotexist123@notexist.notfound",
			},
			expectedError: errors.New("User Not Exist"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			actualRs := friendshipManager.MakeFriend(tc.mockInput)
			assert.Equal(t, tc.expectedError, actualRs)
		})
	}
}

func TestGetUserFriendList(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()

	const numUsers int = 10
	users, ok := insertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	friendshipManager := NewFriendshipManager(tx)
	for i := 1; i < numUsers; i++ {
		assert.NoError(t, friendshipManager.MakeFriend(FrienshipServiceInput{RequestEmail: users[0], TargetEmail: users[i]}))
	}

	expectedListUsers := []string{}
	expectedListUsers = append(expectedListUsers, users[1:]...)

	testCase := []struct {
		scenario       string
		mockInput      user.Users
		expectedResult []string
		expectedError  error
	}{
		{
			scenario:       "Success",
			mockInput:      user.Users{Email: users[0]},
			expectedResult: expectedListUsers,
			expectedError:  nil,
		},
		{
			scenario:       "User not exist",
			mockInput:      user.Users{Email: "usernotexist@notfound.com"},
			expectedResult: nil,
			expectedError:  errors.New("User Not Exist"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			actualRs, err := friendshipManager.GetFriendsList(tc.mockInput)
			if tc.scenario == "Success" {
				assert.Nil(t, err)
				assert.Nil(t, difference(tc.expectedResult, actualRs))
			} else {
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, nil)
			}
		})
	}
}

func TestGetMutualFriendsList(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()

	const numUsers int = 6
	users, ok := insertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	firstUser := users[0]
	secondUser := users[1]

	friendshipManager := NewFriendshipManager(tx)
	for i := 2; i < numUsers; i++ {
		assert.NoError(t, friendshipManager.MakeFriend(FrienshipServiceInput{RequestEmail: firstUser, TargetEmail: users[i]}))
		assert.NoError(t, friendshipManager.MakeFriend(FrienshipServiceInput{RequestEmail: secondUser, TargetEmail: users[i]}))
	}

	expectedMutualFriendsList := []string{}
	expectedMutualFriendsList = append(expectedMutualFriendsList, users[2:]...)

	testCase := []struct {
		scenario       string
		mockInput      FrienshipServiceInput
		expectedResult []string
		expectedError  error
	}{
		{
			scenario: "Success",
			mockInput: FrienshipServiceInput{
				RequestEmail: firstUser,
				TargetEmail:  secondUser,
			},
			expectedResult: expectedMutualFriendsList,
			expectedError:  nil,
		},
		{
			scenario: "User not exist",
			mockInput: FrienshipServiceInput{
				RequestEmail: "usernotexist@notfound.com",
				TargetEmail:  secondUser,
			},
			expectedResult: nil,
			expectedError:  errors.New("User Not Exist"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			actualRs, err := friendshipManager.GetMutualFriendsList(tc.mockInput)
			if tc.scenario == "Success" {
				assert.Nil(t, err)
				assert.Nil(t, difference(tc.expectedResult, actualRs))
			} else {
				assert.Nil(t, actualRs)
				assert.Equal(t, tc.expectedError, err)
			}
		})
	}
}

func TestSubscribe(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()

	const numUsers int = 4
	users, ok := insertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	friendshipManager := NewFriendshipManager(tx)

	testCase := []struct {
		scenario      string
		mockInput     FrienshipServiceInput
		expectedError error
	}{
		{
			scenario: "Success in case both isn't friends",
			mockInput: FrienshipServiceInput{
				RequestEmail: users[0],
				TargetEmail:  users[1],
			},
			expectedError: nil,
		},
		{
			scenario: "Success in case both is friends",
			mockInput: FrienshipServiceInput{
				RequestEmail: users[2],
				TargetEmail:  users[3],
			},
			expectedError: nil,
		},
		{
			scenario: "User not exist",
			mockInput: FrienshipServiceInput{
				RequestEmail: "usernotexist@notfound.com",
				TargetEmail:  users[0],
			},
			expectedError: errors.New("User Not Exist"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			if tc.scenario == "Success in case both is friends" {
				assert.NoError(t, friendshipManager.MakeFriend(tc.mockInput))
			}
			err := friendshipManager.Subscribe(tc.mockInput)
			if tc.scenario == "User not exist" {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBlock(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()

	const numUsers int = 2
	users, ok := insertUsersTest(tx, numUsers)
	assert.Equal(t, true, ok)
	assert.Equal(t, numUsers, len(users))

	firstUser := users[0]
	secondUser := users[1]

	friendshipManager := NewFriendshipManager(tx)

	testCase := []struct {
		scenario      string
		mockInput     FrienshipServiceInput
		expectedError error
	}{
		{
			scenario: "Success in case both is friends",
			mockInput: FrienshipServiceInput{
				RequestEmail: firstUser,
				TargetEmail:  secondUser,
			},
			expectedError: nil,
		},
		{
			scenario: "Success in case both isn't friends",
			mockInput: FrienshipServiceInput{
				RequestEmail: firstUser,
				TargetEmail:  secondUser,
			},
			expectedError: nil,
		},
		{
			scenario: "User not exist",
			mockInput: FrienshipServiceInput{
				RequestEmail: "usernotexist@notfound.com",
				TargetEmail:  secondUser,
			},
			expectedError: errors.New("User Not Exist"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			if tc.scenario == "Success in case both is friends" {
				assert.NoError(t, friendshipManager.MakeFriend(tc.mockInput))
			}
			err := friendshipManager.Block(tc.mockInput)
			if tc.scenario == "User not exist" {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

// ==================================== BEGIN TEST GetUsersReceiveUpdate FUNC =================================
func TestGetUsersReceiveUpdate(t *testing.T) {
	dbconn := utils.CreateConnection()
	tx := dbconn.Begin()
	tx.SavePoint("sp1")
	defer tx.RollbackTo("sp1")

	friendshipManager := NewFriendshipManager(tx)

	// Sender
	sender, ok0 := insertUsersTest(tx, 1)
	assert.Equal(t, true, ok0)
	assert.Equal(t, 1, len(sender))

	// User will be use Make Friend with sender
	const numUsersMakeFriend int = 3
	usersWillMakeFriend, ok1 := insertUsersTest(tx, numUsersMakeFriend)
	assert.Equal(t, true, ok1)
	assert.Equal(t, numUsersMakeFriend, len(usersWillMakeFriend))

	// User will be use subscribe to sender
	const numUsersSubscribe int = 3
	usersSubscribe, ok2 := insertUsersTest(tx, numUsersSubscribe)
	assert.Equal(t, true, ok2)
	assert.Equal(t, numUsersSubscribe, len(usersSubscribe))

	// User mentioned
	const numMentionedUsers int = 2
	mentionedUsers, ok3 := insertUsersTest(tx, numMentionedUsers)
	assert.Equal(t, true, ok3)
	assert.Equal(t, numMentionedUsers, len(mentionedUsers))

	// Make Friend
	for i := 0; i < numUsersMakeFriend; i++ {
		assert.NoError(t, friendshipManager.MakeFriend(FrienshipServiceInput{RequestEmail: usersWillMakeFriend[i], TargetEmail: sender[0]}))
	}

	// Subscribe
	for i := 0; i < numUsersSubscribe; i++ {
		assert.NoError(t, friendshipManager.Subscribe(FrienshipServiceInput{RequestEmail: usersSubscribe[i], TargetEmail: sender[0]}))
	}

	// Expected result
	expectedRs := []string{}
	expectedRs = append(expectedRs, usersWillMakeFriend...)
	expectedRs = append(expectedRs, usersSubscribe...)
	expectedRs = append(expectedRs, mentionedUsers...)

	testCase := []struct {
		scenario               string
		mockSenderInput        string
		mockMentionedUserInput []string
		expectedResult         []string
		expectedError          error
	}{
		{
			scenario:               "Success",
			mockSenderInput:        sender[0],
			mockMentionedUserInput: mentionedUsers,
			expectedResult:         expectedRs,
			expectedError:          nil,
		},
		{
			scenario:               "User not exits",
			mockSenderInput:        "usernotexist@notfound.com",
			mockMentionedUserInput: nil,
			expectedResult:         nil,
			expectedError:          errors.New("User Not Exist"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.scenario, func(t *testing.T) {
			actualRs, err := friendshipManager.GetUsersReceiveUpdate(tc.mockSenderInput, tc.mockMentionedUserInput)
			if tc.scenario == "Success" {
				assert.Nil(t, err)
				assert.Nil(t, difference(tc.expectedResult, actualRs))
			} else {
				assert.Nil(t, actualRs)
				assert.Equal(t, tc.expectedError, err)
			}
		})
	}
}

// InsertUsersTest
func insertUsersTest(tx *gorm.DB, numsUser int) ([]string, bool) {
	listUsers := []string{}
	userManager := user.NewUserManager(tx)
	for i := 0; i < numsUser; i++ {
		email := randomData.Email()
		err := userManager.CreateNewUser(user.Users{Email: email})
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
