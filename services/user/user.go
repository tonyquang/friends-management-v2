package user

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID    uint64 `json:"id" gorm:"column:id; primaryKey"`
	Email string `json:"email" gorm:"column:email; index:unique"`
}
