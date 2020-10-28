package respone

//Using for Retrieve List friends of an user or List common friends of two users
type ResponeListFriends struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   uint64   `json:"count"`
}
