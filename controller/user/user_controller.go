package user

import (
	"friends_management_v2/services/user"
	userService "friends_management_v2/services/user"
	"friends_management_v2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// cmt here
func CreateNewUserHandler(c *gin.Context, service userService.UserService) {
	var ur RequestCreateUser
	if err := c.BindJSON(&ur); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utils.ValidateEmail(ur.Email) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email"})
		return
	}

	rs := service.CreateNewUser(user.Users{Email: ur.Email})

	if rs == nil {
		c.JSON(201, gin.H{"success": true})
		return
	}

	c.JSON(201, rs)
}

func GetListUsersHandler(c *gin.Context, service userService.UserService) {

	rs, err := service.GetListUser()

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, ToListUsers(rs))
}

func ToListUsers(list []string) ResponeListUser {
	listUsers := ResponeListUser{}
	listUsers.ListUsers = append(listUsers.ListUsers, list...)
	listUsers.Count = uint64(len(list))
	return listUsers
}
