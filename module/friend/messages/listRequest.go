package friend

//ListRequest : struct to hold request list of connections
type ListRequest struct {
	Email string `json:"email" binding:"required,email"`
}
