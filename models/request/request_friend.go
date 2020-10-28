package request

//Using for Request Add Friend and Retrieve the common friends
type RequestFriend struct {
	Friends []string `json:"friends"`
}
