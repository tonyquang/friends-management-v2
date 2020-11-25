package controller

import (
	"fmt"
	"net/http"

	friendshipController "friends_management_v2/controller/friendship"
	userController "friends_management_v2/controller/user"
	migration "friends_management_v2/migrations"
	service "friends_management_v2/services"
	friendshipService "friends_management_v2/services/friendship"
	userService "friends_management_v2/services/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"gorm.io/gorm"
)

//Setup Manager, Migration and Routes
func Setup(db *gorm.DB) http.Handler {
	friendshipService := friendshipService.NewFriendshipManager(db)
	userService := userService.NewUserManager(db)

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)

	migration.InitMigration(db)
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	//url := ginSwagger.URL("http://localhost:3000/docs/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		friendshipController.GetFriendsListController(c, friendshipService)
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

	r.POST("/get-list-users-receive-update", func(c *gin.Context) {
		friendshipController.GetUsersReceiveUpdateController(c, friendshipService)
	})
	return r
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
