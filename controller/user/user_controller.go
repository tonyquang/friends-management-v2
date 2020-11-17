package user

import (
	"net/http"

	httpRes "friends_management_v2/controller/common_respone"
	"friends_management_v2/services/user"
	userService "friends_management_v2/services/user"
	"friends_management_v2/utils"

	"github.com/gin-gonic/gin"
)

// Paths Information

// CreateNewUserController godoc
// @Summary Create A New User
// @Description Create A New User
// @Tags User
// @Consume json
// @Param email body RequestCreateUser true "RequestCreateUser"
// @Produce  json
// @Success 201 {object} httpRes.HTTPSuccess
// @Failure 400 {object} httpRes.HTTPError
// @Router /create-user [post]
func CreateNewUserController(c *gin.Context, service userService.UserService) {
	var ur RequestCreateUser
	if err := c.BindJSON(&ur); err != nil {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "BindJson Error, cause body request invalid"})
		return
	}

	if utils.ValidateEmail(ur.Email) == false {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Invalid Email"})
		return
	}

	rs := service.CreateNewUser(user.Users{Email: ur.Email})

	if rs == nil {
		c.JSON(201, httpRes.HTTPSuccess{Success: true})
		return
	}

	c.JSON(400, httpRes.HTTPError{Message: rs.Error()})
}

// GetListUsersController godoc
// @Summary List users
// @Description Get list users
// @Tags User
// @Produce  json
// @Success 200 {object} ResponeListUser
// @Failure 500 {object} httpRes.HTTPError
// @Router /list-users [get]
func GetListUsersController(c *gin.Context, service userService.UserService) {

	rs, err := service.GetListUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, HTTPError{Message: err.Error()})
		return
	}

	c.JSON(200, toListUsers(rs))
}

func toListUsers(list []string) ResponeListUser {
	listUsers := ResponeListUser{}
	listUsers.ListUsers = append(listUsers.ListUsers, list...)
	listUsers.Count = uint64(len(list))
	return listUsers
}
