package user

import (
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (_m *UserRepoMock) CheckUserExist(emailAddress []string) (bool, error) {
	args := _m.Called(emailAddress)
	return args.Bool(0), args.Error(1)
}
