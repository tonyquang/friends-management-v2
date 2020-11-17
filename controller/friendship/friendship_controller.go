package friendship

import (
	"net/http"

	httpRes "friends_management_v2/controller/common_respone"
	"friends_management_v2/services/friendship"
	"friends_management_v2/services/user"
	"friends_management_v2/utils"

	"github.com/gin-gonic/gin"
)

// Paths Information

// CreateNewUserController godoc
// @Summary Make Friend Connection
// @Description Create a friend connection between two email addresses.
// @Tags Friendship
// @Consume json
// @Param friends body RequestFriend true "RequestCreateUser"
// @Produce  json
// @Success 201 {object} httpRes.HTTPSuccess
// @Failure 400 {object} httpRes.HTTPError
// @Router /add-friends [post]
func MakeFriendController(c *gin.Context, service friendship.FrienshipServices) {
	var reqFriend RequestFriend

	if err := c.BindJSON(&reqFriend); err != nil {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "BindJson Error, cause body request invalid"})
		return
	}

	if len(reqFriend.Friends) != 2 || reqFriend.Friends[0] == reqFriend.Friends[1] {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Request Invalid"})
		return
	}

	firstUser := reqFriend.Friends[0]
	secondUser := reqFriend.Friends[1]

	if utils.ValidateEmail(firstUser) == false || utils.ValidateEmail(secondUser) == false {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Email Invalid Format"})
		return
	}

	rs := service.MakeFriend(friendship.FrienshipServiceInput{RequestEmail: firstUser, TargetEmail: secondUser})

	if rs == nil {
		c.JSON(201, httpRes.HTTPSuccess{Success: true})
		return
	}

	c.JSON(400, httpRes.HTTPError{Message: rs.Error()})
}

// GetFriendsListController godoc
// @Summary Get Friends List
// @Description Retrieve the friends list for an email address.
// @Tags Friendship
// @Consume json
// @Param email body RequestListFriends true "RequestListFriends"
// @Produce  json
// @Success 200 {object} ResponeListFriends
// @Failure 400 {object} httpRes.HTTPError
// @Router /get-list-friends [post]
func GetFriendsListController(c *gin.Context, service friendship.FrienshipServices) {
	email := RequestListFriends{}

	if err := c.BindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "BindJson Error, cause body request invalid"})
		return
	}

	if utils.ValidateEmail(email.Mail) == false {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Email Invalid Format"})
		return
	}

	rs, err := service.GetFriendsList(user.Users{Email: email.Mail})

	if err != nil {
		c.JSON(400, httpRes.HTTPError{Message: err.Error()})
		return
	}

	c.JSON(200, toListFriendsStruct(rs))
}

// GetMutualFriendsController godoc
// @Summary Get Mutual Friends List
// @Description Retrieve the common friends list between two email addresses.
// @Tags Friendship
// @Consume json
// @Param friends body RequestFriend true "RequestFriend"
// @Produce  json
// @Success 200 {object} ResponeListFriends
// @Failure 400 {object} httpRes.HTTPError
// @Router /get-mutual-list-friends [post]
func GetMutualFriendsController(c *gin.Context, service friendship.FrienshipServices) {
	reqFriend := RequestFriend{}

	if err := c.BindJSON(&reqFriend); err != nil {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "BindJson Error, cause body request invalid"})
		return
	}

	if len(reqFriend.Friends) != 2 || reqFriend.Friends[0] == reqFriend.Friends[1] {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Request Invalid"})
		return
	}

	firstUser := reqFriend.Friends[0]
	secondUser := reqFriend.Friends[1]

	if utils.ValidateEmail(firstUser) == false || utils.ValidateEmail(secondUser) == false {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Email Invalid Format"})
		return
	}

	rs, err := service.GetMutualFriendsList(friendship.FrienshipServiceInput{RequestEmail: firstUser, TargetEmail: secondUser})

	if err != nil {
		c.JSON(400, httpRes.HTTPError{Message: err.Error()})
		return
	}

	c.JSON(200, toListFriendsStruct(rs))
}

// SubscribeController godoc
// @Summary Subscribe update an user
// @Description Subscribe to updates from an email address.
// @Tags Friendship
// @Consume json
// @Param request body RequestUpdate true "Requestor and Target to subscribe update"
// @Produce  json
// @Success 201 {object} httpRes.HTTPSuccess
// @Failure 400 {object} httpRes.HTTPError
// @Router /subscribe [post]
func SubscribeController(c *gin.Context, service friendship.FrienshipServices) {
	reqSubscribe := RequestUpdate{}

	if err := c.BindJSON(&reqSubscribe); err != nil {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "BindJson Error, cause body request invalid"})
		return
	}

	if reqSubscribe.Requestor == reqSubscribe.Target {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Request Invalid"})
		return
	}

	firstUser := reqSubscribe.Requestor
	secondUser := reqSubscribe.Target

	if utils.ValidateEmail(firstUser) == false || utils.ValidateEmail(secondUser) == false {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Email Invalid Format"})
		return
	}

	rs := service.Subscribe(friendship.FrienshipServiceInput{RequestEmail: firstUser, TargetEmail: secondUser})

	if rs != nil {
		c.JSON(400, httpRes.HTTPError{Message: rs.Error()})
		return
	}

	c.JSON(201, httpRes.HTTPSuccess{Success: true})
}

// BlockController godoc
// @Summary Block Subscribe to update an user
// @Description Block updates from an email address.
// @Tags Friendship
// @Consume json
// @Param request body RequestUpdate true "Requestor and Target to block update"
// @Produce  json
// @Success 201 {object} httpRes.HTTPSuccess
// @Failure 400 {object} httpRes.HTTPError
// @Router /block [post]
func BlockController(c *gin.Context, service friendship.FrienshipServices) {
	reqSubscribe := RequestUpdate{}

	if err := c.BindJSON(&reqSubscribe); err != nil {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "BindJson Error, cause body request invalid"})
		return
	}

	if reqSubscribe.Requestor == reqSubscribe.Target {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Request Invalid"})
		return
	}

	firstUser := reqSubscribe.Requestor
	secondUser := reqSubscribe.Target

	if utils.ValidateEmail(firstUser) == false || utils.ValidateEmail(secondUser) == false {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Email Invalid Format"})
		return
	}

	rs := service.Block(friendship.FrienshipServiceInput{RequestEmail: firstUser, TargetEmail: secondUser})

	if rs != nil {
		c.JSON(400, httpRes.HTTPError{Message: rs.Error()})
		return
	}

	c.JSON(201, httpRes.HTTPSuccess{Success: true})
}

// GetUsersReceiveUpdateController godoc
// @Summary Get Users Receive Update
// @Description Retrieve all email addresses that can receive updates from an email address.
// @Tags Friendship
// @Consume json
// @Param request body RequestReceiveUpdate true "Sender and Text"
// @Produce  json
// @Success 200 {object} ResponeReceiveUpdate
// @Failure 400 {object} httpRes.HTTPError
// @Router /get-list-users-receive-update [post]
func GetUsersReceiveUpdateController(c *gin.Context, service friendship.FrienshipServices) {
	reqRecvUpdate := RequestReceiveUpdate{}

	if err := c.BindJSON(&reqRecvUpdate); err != nil {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "BindJson Error, cause body request invalid"})
		return
	}

	if utils.ValidateEmail(reqRecvUpdate.Sender) == false {
		c.JSON(http.StatusBadRequest, httpRes.HTTPError{Message: "Email Invalid Format"})
		return
	}

	// rename
	mentionedUsers := utils.ExtractMentionEmail(reqRecvUpdate.Text)

	rs, err := service.GetUsersReceiveUpdate(reqRecvUpdate.Sender, mentionedUsers)

	if err != nil {
		c.JSON(400, httpRes.HTTPError{Message: err.Error()})
		return
	}

	c.JSON(200, toUsersCanReceiveUpdate(removeDuplicates(rs)))
}

func toListFriendsStruct(list []string) ResponeListFriends {
	listFriendsRespone := ResponeListFriends{}
	listFriendsRespone.Count = uint64(len(list))
	listFriendsRespone.Success = true
	listFriendsRespone.Friends = append(listFriendsRespone.Friends, list...)
	return listFriendsRespone
}

func toUsersCanReceiveUpdate(list []string) ResponeReceiveUpdate {
	listUsersRecvUpdate := ResponeReceiveUpdate{}
	listUsersRecvUpdate.Success = true
	listUsersRecvUpdate.Recipients = append(listUsersRecvUpdate.Recipients, list...)
	return listUsersRecvUpdate
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
