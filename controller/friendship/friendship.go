package friendship

//Using for Retrieve List friends of an user or List common friends of two users
type ResponeListFriends struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   uint64   `json:"count"`
}

type ResponeReciveUpdate struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

//Using for Request Add Friend and Retrieve the common friends
type RequestFriend struct {
	Friends []string `json:"friends"`
}

type RequestReciveUpdate struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

//Using for request Subscribe or Block Update
type RequestSubscribe struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
