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
	ID           uint64 `json:"id"`
	FirstUser    string `json:"first_user"`
	SecondUser   string `json:"second_user"`
	IsFriend     bool   `json:"is_friend"`
	UpdateStatus bool   `json:"update_status"`
}
