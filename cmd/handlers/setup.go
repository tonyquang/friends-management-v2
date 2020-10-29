package handlers

import (
	"friends_management_v2/migration"
	"friends_management_v2/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Setup Manager, Migration and Routes
func Setup(db *gorm.DB) http.Handler {
	friendshipServices := services.NewManager(db)
	migration.InitMigration(db)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/create-user", func(c *gin.Context) {
		CreateNewUserHandler(c, friendshipServices)
	})

	r.POST("/create-friend-connection", func(c *gin.Context) {
		CreateNewFriendConnection(c, friendshipServices)
	})

	return r
}
