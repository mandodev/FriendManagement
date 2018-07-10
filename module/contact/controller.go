package contact

//Controller for contact management
type Controller struct {
}

//NewController : function to instantiate new controller
func NewController() (*Controller, error) {
	return &Controller{}, nil
}
