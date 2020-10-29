package repository

import (
	"friends_management_v2/models/db_models"
	"friends_management_v2/models/respone"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
	return FinalReturn(rs, "Insert Data To Table User Erorr: ")

}

//create a friend connection between two email addresses.
func InsertNewFriendConnection(dbconn *gorm.DB, firstUser, secondUser string) *respone.ResponeError {
	IsConnection, err := CheckConnection(dbconn, firstUser, secondUser)

	if err != nil {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusInternalServerError, Description: "Check Connection Users Error: " + err.Error()}
	}

	if IsConnection == true {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusBadRequest, Description: "Connection is already"}
	}

	listUsers := [2]string{firstUser, secondUser}

	tx := dbconn.Session(&gorm.Session{
		WithConditions: true,
		Logger:         dbconn.Logger.LogMode(logger.Info),
	})

	for i := 0; i < 2; i++ {

		ok1, err1 := CheckUserExist(tx, listUsers[i])

		if err1 != nil {
			return &respone.ResponeError{Success: false, StatusCode: http.StatusInternalServerError, Description: "Check Is Exist User Error: " + err.Error()}
		}

		if ok1 == false {
			return &respone.ResponeError{Success: false, StatusCode: http.StatusBadRequest, Description: "User not exist!"}
		}
	}

	rs := dbconn.Create(&db_models.Friendship{FirstUser: firstUser, SecondUser: secondUser, IsFriend: true, UpdateStatus: false})
	return FinalReturn(rs, "Insert New Connection Error: ")
}
