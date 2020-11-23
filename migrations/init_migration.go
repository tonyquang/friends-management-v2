package migration

import (
	"friends_management_v2/services/friendship"
	"friends_management_v2/services/user"

	"gorm.io/gorm"
)

func InitMigration(dbconn *gorm.DB) {
<<<<<<< HEAD
	dbconn.AutoMigrate(&user.Users{})
	dbconn.AutoMigrate(&friendship.Friendship{})
=======

	if oke := dbconn.Migrator().HasTable(&user.Users{}); !oke {
		dbconn.AutoMigrate(&user.Users{})
	}

	if oke := dbconn.Migrator().HasTable(&friendship.Friendship{}); !oke {
		dbconn.AutoMigrate(&friendship.Friendship{})
	}
>>>>>>> DONE-API
}
