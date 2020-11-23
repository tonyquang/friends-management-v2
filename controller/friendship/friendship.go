package friendship

// Using for Retrieve List friends of an user or List common friends of two users
type ResponeListFriends struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   uint     `json:"count"`
}

type ResponeReceiveUpdate struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

// Using for Request Add Friend and Retrieve the common friends
type RequestFriend struct {
	Friends []string `json:"friends" binding:"required"`
}

type RequestListFriends struct {
	Mail string `json:"email" binding:"required"`
}

type RequestReceiveUpdate struct {
	Sender string `json:"sender" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

// Using for request Subscribe or Block Update
type RequestUpdate struct {
	Requestor string `json:"requestor" binding:"required"`
	Target    string `json:"target" binding:"required"`
}
