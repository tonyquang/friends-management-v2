package friendship

import (
	"github.com/stretchr/testify/mock"
)

type FriendshipRepoMock struct {
	mock.Mock
}

func (_m *FriendshipRepoMock) checkFriendship(firstUser, secondUser string) (*Friendship, error) {
	args := _m.Called(firstUser, secondUser)
	return args.Get(0).(*Friendship), args.Error(1)
}

func (_m *FriendshipRepoMock) execMakeFriend(requestor string, target string) error {
	args := _m.Called(requestor, target)
	return args.Error(0)
}

func (_m *FriendshipRepoMock) execUpdateMakeFriend(input []string) error {
	args := _m.Called(input)
	return args.Error(0)
}

func (_m *FriendshipRepoMock) checkUserExist(listUsers []string) (bool, error) {
	args := _m.Called(listUsers)
	return args.Bool(0), args.Error(1)
}
