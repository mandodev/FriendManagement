package notification

//NotificationRequest : struct for subscribe request
type NotificationRequest struct {
	Requestor string `json:"requestor" binding:"required" validator:"email"`
	Target    string `json:"target" binding:"required" validator:"email"`
}
