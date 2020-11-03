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
		userController.GetListUsersController(c, userService)
	})

	r.POST("/create-user", func(c *gin.Context) {
		userController.CreateNewUserController(c, userService)
	})

	r.POST("/create-friend-connection", func(c *gin.Context) {
		friendshipController.AddFriendController(c, friendshipService)
	})

	r.POST("/get-list-friends", func(c *gin.Context) {
		friendshipController.GetFriendsListAnUserController(c, friendshipService)
	})

	r.POST("/get-mutual-list-friends", func(c *gin.Context) {
		friendshipController.GetMutualFriendsListController(c, friendshipService)
	})

	r.POST("/subscribe-update", func(c *gin.Context) {
		friendshipController.SubscribeUpdateController(c, friendshipService)
	})
	return r
}
