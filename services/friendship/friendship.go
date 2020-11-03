package friendship

import (
	"gorm.io/gorm"
)

type ServiceFrienshipInput struct {
	First_user  string
	Second_user string
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
// -Case 1: equal 0 when FirstUser and SecondUser Make A new friend Connection and Both not subcribe update each other
// -Case 2:
// 			+ equal 1 when FirstUser subscribe update to SecondUser
//			+ equal -1 when FirstUser block update to SecondUser
// -Case 3:
// 			+ equal 2 when SecondUser subscribe update to FirstUser
//			+ equal -2 when SecondUser block update to FirstUser
// -Case 4: equal 3 when both subscribe update togerther