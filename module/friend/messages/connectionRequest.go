package friend

//ConnectionRequest : message structure for connection request
type ConnectionRequest struct {
	Friends []string `json:"friends" binding:"required,len=2"`
}
