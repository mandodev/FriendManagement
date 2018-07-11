package notification

//Controller : Notification Controller
type Controller struct {
	notificationService *Service
}

//NewController : function to instantiate new controller
func NewController(service *Service) (*Controller, error) {
	return &Controller{notificationService: service}, nil
}
