package migration

import (
	"friends_management_v2/services/friendship"
	"friends_management_v2/services/user"

	"gorm.io/gorm"
)

func InitMigration(dbconn *gorm.DB) {
	dbconn.AutoMigrate(&user.Users{})
	dbconn.AutoMigrate(&friendship.Friendship{})
}
