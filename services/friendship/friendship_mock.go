package friendship

import (
	"friends_management_v2/services/user"

	"github.com/stretchr/testify/mock"
)

type FrienshipMockService struct {
	mock.Mock
}

func (_m *FrienshipMockService) MakeFriend(input ServiceFrienshipInput) error {
	args := _m.Called(input)
	output := args.Error(0)
	return output
}

func (_m *FrienshipMockService) GetUserFriendList(ur user.Users) ([]string, error) {
	args := _m.Called(ur)
	return args.Get(0).([]string), args.Error(1)
}
func (_m *FrienshipMockService) GetMutualFriendsList(input ServiceFrienshipInput) ([]string, error) {
	args := _m.Called(input)
	return args.Get(0).([]string), args.Error(1)
}
func (_m *FrienshipMockService) Subcribe(input ServiceFrienshipInput) error {
	args := _m.Called(input)
	return args.Error(1)
}
func (_m *FrienshipMockService) Block(input ServiceFrienshipInput) error {
	args := _m.Called(input)
	return args.Error(1)
}
func (_m *FrienshipMockService) GetUsersRecevieUpdate(sender string, mentionedUsers []string) ([]string, error) {
	args := _m.Called(sender, mentionedUsers)
	return args.Get(0).([]string), args.Error(1)
}
