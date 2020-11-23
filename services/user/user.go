package user

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
<<<<<<< HEAD
	ID    uint64 `json:"id"`
	Email string `json:"email"`
=======
	ID    uint64 `json:"id" gorm:"column:id; primaryKey"`
	Email string `json:"email" gorm:"column:email; index:unique"`
>>>>>>> DONE-API
}
