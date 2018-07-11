package notification

//SubscribeRequest : struct for subscribe request
type SubscribeRequest struct {
	Requestor string `json:"requestor" binding:"required" validator:"email"`
	Target    string `json:"target" binding:"required" validator:"email"`
}
