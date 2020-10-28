package request

//Using for request Subscribe or Block Update
type RequestUpdate struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
