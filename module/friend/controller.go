package friend

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"

	messages "github.com/FriendManagement/module/friend/messages"
	"github.com/gin-gonic/gin"
)

//Controller for contact management
type Controller struct {
	connectionService *Service
}

//NewController : function to instantiate new controller
func NewController(service *Service) (*Controller, error) {
	return &Controller{connectionService: service}, nil
}

//CreateConnection : Create connection between two email
func (c *Controller) CreateConnection(ctx *gin.Context) {
	var request messages.ConnectionRequest
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

	result, err := c.connectionService.CreateConnection(request.Friends)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": result, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": result})
}

//List : List user's connections
func (c *Controller) List(ctx *gin.Context) {
	var request messages.ListRequest
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

	result, err := c.connectionService.ConnectionList(request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": result, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "friends": result, "count": len(result)})

}

//Common : List common connection between two user
func (c *Controller) Common(ctx *gin.Context) {
	var request messages.CommonRequest
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

	result, err := c.connectionService.CommonConnection(request.Friends)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": result, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "friends": result, "count": len(result)})
}
