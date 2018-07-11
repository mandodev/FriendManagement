package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Controller : Notification Controller
type Controller struct {
	notificationService *Service
}

//NewController : function to instantiate new controller
func NewController(service *Service) (*Controller, error) {
	return &Controller{notificationService: service}, nil
}

//Subscribe : function that allow user to subscribe update based on email address
func (c *Controller) Subscribe(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
