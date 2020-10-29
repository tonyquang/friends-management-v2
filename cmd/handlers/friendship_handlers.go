package handlers

import (
	"friends_management_v2/models/request"
	"friends_management_v2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewUserHandler(c *gin.Context, service services.Services) {
	var user request.RequestCreateUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rs := service.CreateNewUser(user)

	if rs == nil {
		c.JSON(201, gin.H{"success": true})
		return
	}

	c.JSON(rs.StatusCode, rs)
	// res, err := service.AddFriend(friends)

	// utils.Respone(res, err, c)
}
