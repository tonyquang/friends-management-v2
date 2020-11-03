package friendship

import (
	"friends_management_v2/services/friendship"
	"friends_management_v2/services/user"
	"friends_management_v2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFriendController(c *gin.Context, service friendship.FrienshipServices) {
	var reqFriend RequestFriend

	if err := c.BindJSON(&reqFriend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(reqFriend.Friends) != 2 || reqFriend.Friends[0] == reqFriend.Friends[1] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request Invalid"})
		return
	}

	firstUser := reqFriend.Friends[0]
	secondUser := reqFriend.Friends[1]

	if utils.ValidateEmail(firstUser) == false || utils.ValidateEmail(secondUser) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Invalid Format"})
		return
	}

	rs := service.AddFriends(friendship.ServiceFrienshipInput{First_user: firstUser, Second_user: secondUser})

	if rs == nil {
		c.JSON(201, gin.H{"success": true})
		return
	}

	c.JSON(400, gin.H{"error": rs.Error()})
}

func GetFriendsListAnUserController(c *gin.Context, service friendship.FrienshipServices) {
	email := struct {
		Mail string `json:"email"`
	}{}

	if err := c.BindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utils.ValidateEmail(email.Mail) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Invalid Format"})
		return
	}

	rs, err := service.GetFriendListOfAnUser(user.Users{Email: email.Mail})

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ToListFriendsStruct(rs))
}

func GetMutualFriendsListController(c *gin.Context, service friendship.FrienshipServices) {
	reqFriend := RequestFriend{}

	if err := c.BindJSON(&reqFriend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(reqFriend.Friends) != 2 || reqFriend.Friends[0] == reqFriend.Friends[1] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request Invalid"})
		return
	}

	firstUser := reqFriend.Friends[0]
	secondUser := reqFriend.Friends[1]

	if utils.ValidateEmail(firstUser) == false || utils.ValidateEmail(secondUser) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Invalid Format"})
		return
	}

	rs, err := service.GetMutualFriendsList(friendship.ServiceFrienshipInput{First_user: firstUser, Second_user: secondUser})

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ToListFriendsStruct(rs))
}

func SubscribeUpdateController(c *gin.Context, service friendship.FrienshipServices) {
	reqSubscribe := RequestSubscribe{}

	if err := c.BindJSON(&reqSubscribe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if reqSubscribe.Requestor == reqSubscribe.Target {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request Invalid"})
		return
	}

	firstUser := reqSubscribe.Requestor
	secondUser := reqSubscribe.Target

	if utils.ValidateEmail(firstUser) == false || utils.ValidateEmail(secondUser) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Invalid Format"})
		return
	}

	rs := service.SubscribeUpdate(friendship.ServiceFrienshipInput{First_user: firstUser, Second_user: secondUser})

	if rs != nil {
		c.JSON(400, gin.H{"error": rs.Error()})
		return
	}

	c.JSON(200, gin.H{"success": true})
}

func ToListFriendsStruct(list []string) ResponeListFriends {
	listFriendsRespone := ResponeListFriends{}
	listFriendsRespone.Count = uint64(len(list))
	listFriendsRespone.Success = true
	listFriendsRespone.Friends = append(listFriendsRespone.Friends, list...)
	return listFriendsRespone
}
