package notification

import (
	"fmt"
	"net/http"

	messages "github.com/FriendManagement/module/notification/messages"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
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
	var request messages.SubscribeRequest
	var errors []string

	if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if ok {
			for _, v := range ve {
				msg := fmt.Sprintf("%s is %s", v.Field, v.Tag)
				if v.Tag == "len" {
					msg = fmt.Sprintf("%s %s should be %s", v.Field, v.Tag, v.Param)
				}
				errors = append(errors, msg)
			}
		} else {
			errors = append(errors, err.Error())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "errors": errors})
		return
	}
	result, err := c.notificationService.Subscribe(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": result, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
