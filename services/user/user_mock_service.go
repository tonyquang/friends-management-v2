package user

import (
	"github.com/stretchr/testify/mock"
)

type UserMockService struct {
	mock.Mock
}

func (_m *UserMockService) CreateNewUser(userMail Users) error {
	args := _m.Called(userMail)
	return args.Error(0)
}
func (_m *UserMockService) GetListUser() ([]string, error) {
	args := _m.Called()
	return args.Get(0).([]string), args.Error(1)
}
