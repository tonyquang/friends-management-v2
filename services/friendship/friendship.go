package friendship

import (
	"gorm.io/gorm"
)

type FrienshipServiceInput struct {
	RequestEmail string
	TargetEmail  string
}

type Friendship struct {
	gorm.Model
	ID           uint64 `json:"id" gorm:"column:id"`
	FirstUser    string `json:"first_user" gorm:"column:first_user"`
	SecondUser   string `json:"second_user" gorm:"column:second_user"`
	IsFriend     bool   `json:"is_friend" gorm:"column:is_friend"`
	UpdateStatus int    `json:"update_status" gorm:"column:update_status"`
}

// UpdateStatus Mean
// -Case 1:
//  + equal 0 when FirstUser and SecondUser block update together
// A <--------------x------------- B
// A -------------x--------------> B

// -Case 2:														 ---
// 	+ equal 1 when FirstUser subscribe update to SecondUser		   |
//	+ and SecondUser block update to FirstUser					   |
// A <------------x--------------- B							   |
// A ----------------------------> B							   |
//																   |
// -Case 3:															\ 	Case 2 and Case 3 using in case when Make Friend
// 	+ equal 2 when SecondUser subscribe update to FirstUser			/
//	+ and FirstUser block update to SecondUser					   |
// A <---------------------------- B							   |
// A -------------x--------------> B							   |
//																 ---
// -Case 4:
//  + equal 3 when both subscribe update togerther
// A <---------------------------- B
// A ----------------------------> B
