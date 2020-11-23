package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserService interface {
	CreateNewUser(userMail Users) error
	GetListUser() ([]string, error)
}

type UserRepo interface {
}

type UserManager struct {
	dbconn *gorm.DB
}

func NewUserManager(dbconn *gorm.DB) *UserManager {
	return &UserManager{
		dbconn: dbconn,
	}
}

func (m *UserManager) CreateNewUser(userMail Users) error {

	emailAddress := userMail.Email

	IsExist, err := m.CheckUserExist([]string{emailAddress})
	if err != nil {
		return err
	}

	if IsExist == true {
		return errors.New("User is already exists!")
	}

	rs := m.dbconn.Create(&userMail)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (m *UserManager) GetListUser() ([]string, error) {
	listUser := []string{}

	rs := m.dbconn.Select("email").Find(&Users{}).Scan(&listUser)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return listUser, nil
}

func (m *UserManager) CheckUserExist(emailAddress []string) (bool, error) {

	var count int

	rs := m.dbconn.Select("COUNT(*)").Where("email IN ?", emailAddress).Find(&Users{}).Scan(&count)

	if rs.Error != nil {
		return false, rs.Error
	}
	if count == len(emailAddress) {
		return true, nil
	} else {
		return false, nil
	}
}
