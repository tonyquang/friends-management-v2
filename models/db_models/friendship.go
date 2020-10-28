package db_models

import (
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	ID           uint64 `json:"id"`
	FirstUser    string `json:"first_user"`
	SecondUser   string `json:"second_user"`
	IsFriend     bool   `json:"is_friend"`
	UpdateStatus bool   `json:"update_status"`
}
