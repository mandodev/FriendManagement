package shared

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

//SetupRouter : function that return registered end point
func SetupRouter() *gin.Engine {
	router := gin.New()

	//middleware setup
	router.Use(ginglog.Logger(5), gin.Recovery())

	//diagnostic endpoint
	web := router.Group("v1/api")
	{
		web.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Name":       "Friend Management",
				"message":    "OK",
				"serverTime": time.Now().UTC(),
				"version":    "0.0.1",
			})
		})
	}

	return router
}
