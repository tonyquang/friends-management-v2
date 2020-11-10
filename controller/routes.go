package controller

import (
	migration "friends_management_v2/migrations"
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
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/list-users", func(c *gin.Context) {
		userController.GetListUsersController(c, userService)
	})

	r.POST("/create-user", func(c *gin.Context) {
		userController.CreateNewUserController(c, userService)
	})

	r.POST("/add-friends", func(c *gin.Context) {
		friendshipController.MakeFriendController(c, friendshipService)
	})

	r.POST("/get-list-friends", func(c *gin.Context) {
		friendshipController.GetFriendList(c, friendshipService)
	})

	r.POST("/get-mutual-list-friends", func(c *gin.Context) {
		friendshipController.GetMutualFriendsController(c, friendshipService)
	})

	r.POST("/subscribe", func(c *gin.Context) {
		friendshipController.SubscribeController(c, friendshipService)
	})

	r.POST("/block", func(c *gin.Context) {
		friendshipController.BlockController(c, friendshipService)
	})

	r.POST("/get-list-users-revcive-update", func(c *gin.Context) {
		friendshipController.GetUsersRecvUpdateController(c, friendshipService)
	})
	return r
}
