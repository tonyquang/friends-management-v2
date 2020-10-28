package db_models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
