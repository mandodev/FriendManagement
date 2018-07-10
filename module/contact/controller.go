package contact

import (
	"github.com/gin-gonic/gin"
)

//Controller for contact management
type Controller struct {
}

//NewController : function to instantiate new controller
func NewController() (*Controller, error) {
	return &Controller{}, nil
}

//CreateConnection : Create connection between two email
func (c *Controller) CreateConnection(ctx *gin.Context) {

}
