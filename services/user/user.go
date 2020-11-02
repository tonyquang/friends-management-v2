package user

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}
