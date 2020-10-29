package repository

import (
	"fmt"
	"friends_management_v2/models/db_models"
	"friends_management_v2/models/respone"
	"net/http"

	"gorm.io/gorm"
)

//Check Email User Is Exist In Table User
func CheckUserExist(dbconn *gorm.DB, emailAddress string) (bool, error) {
	fmt.Println(emailAddress)
	rs := dbconn.Where("email = ?", emailAddress).Find(&db_models.Users{})

	if rs.Error != nil {
		return false, rs.Error
	}
	fmt.Println(rs.RowsAffected)
	if rs.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}

}

//Check Connection Between Two User
func CheckConnection(dbconn *gorm.DB, firstUser, secondUser string) (bool, error) {
	rs := dbconn.Where("first_user IN ? AND second_user IN ?", []string{firstUser, secondUser}, []string{firstUser, secondUser}).Find(&db_models.Friendship{})
	if rs.Error != nil {
		return false, rs.Error
	}
	if rs.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// Final return repo
func FinalReturn(tx *gorm.DB, messageErr string) *respone.ResponeError {
	if tx.Error != nil {
		return &respone.ResponeError{Success: false, StatusCode: http.StatusInternalServerError, Description: messageErr + tx.Error.Error()}
	}

	return nil
}
