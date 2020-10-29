package migration

import (
	"friends_management_v2/models/db_models"

	"gorm.io/gorm"
)

func InitMigration(dbconn *gorm.DB) {
	dbconn.AutoMigrate(&db_models.Users{})
	dbconn.AutoMigrate(&db_models.Friendship{})
}
