package friend

//CommonRequest : request struct to get common connection between two user
type CommonRequest struct {
	Friends []string `json:"friends" binding:"required,len=2"`
}
