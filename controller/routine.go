package controller

import (
	"friends_management_v2/migration"
	friendshipService "friends_management_v2/services/friendship"
	userService "friends_management_v2/services/user"
	"net/http"

	friendshipController "friends_management_v2/controller/friendship"

	userController "friends_management_v2/controller/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Setup Manager, Migration and Routes
func Setup(db *gorm.DB) http.Handler {
	friendshipService := friendshipService.NewFriendshipManager(db)
	userService := userService.NewUserManager(db)
	migration.InitMigration(db)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/list-users", func(c *gin.Context) {
		userController.GetListUsersHandler(c, userService)
	})

	r.POST("/create-user", func(c *gin.Context) {
		userController.CreateNewUserHandler(c, userService)
	})

	r.POST("/create-friend-connection", func(c *gin.Context) {
		friendshipController.CreateNewFriendConnectionHandler(c, friendshipService)
	})

	r.POST("/get-list-friends", func(c *gin.Context) {
		friendshipController.GetFriendsListAnUserHandler(c, friendshipService)
	})

	r.POST("/get-mutual-list-friends", func(c *gin.Context) {
		friendshipController.GetMutualFriendsList(c, friendshipService)
	})
	return r
}
