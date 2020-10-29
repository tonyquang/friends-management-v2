package repository

import (
	"friends_management_v2/models/db_models"
	"friends_management_v2/models/respone"
	"net/http"

	"gorm.io/gorm"
)

var modelUser db_models.Users

//Insert a new user to Users Table
func InsertNewUser(dbconn *gorm.DB, user db_models.Users) *respone.ResponeError {
	IsExist, err := CheckUserExist(dbconn, user.Email)
	if err != nil {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusInternalServerError, Description: "Check Is Exist User Error: " + err.Error()}
	}

	if IsExist == true {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusBadRequest, Description: "User is already!"}
	}

	rs := dbconn.Create(&user)
	if rs.Error != nil {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusInternalServerError, Description: "Insert Data To Table User Erorr: " + rs.Error.Error()}
	}

	return nil
}

//Check Email User Is Exist In Table User
func CheckUserExist(dbconn *gorm.DB, emailAddress string) (bool, error) {
	rs := dbconn.Where("email = ?", emailAddress).Find(&modelUser)

	if rs.Error != nil {
		return false, rs.Error
	}

	if rs.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}

}
