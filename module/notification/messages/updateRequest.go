package notification

//UpdateRequest : struct for subscribe request
type UpdateRequest struct {
	Sender string `json:"sender" binding:"required" validator:"email"`
	Text   string `json:"text" binding:"required"`
}
