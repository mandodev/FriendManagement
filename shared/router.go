package shared

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/golang/glog"

	"github.com/FriendManagement/module/friend"
	"github.com/FriendManagement/shared/config"
	"github.com/FriendManagement/shared/data"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

//Router : Instance struct for router model
type Router struct {
	database         *gorm.DB
	config           *config.Configuration
	friendController friend.Controller
	friendService    friend.Service
}

//NewRouter : Instantiate new Router
func NewRouter(configuration config.Configuration) *Router {
	cfg, _ := config.New("../../shared/config/")

	dbInstance, err := data.NewDbFactory(cfg)
	dbase, err := dbInstance.DBConnection()
	fService, err := friend.NewService(dbase)

	fController, err := friend.NewController(fService)
	if err != nil {
		glog.Fatalf("Fatal Error on create friend Controller : %s", err.Error())
	}

	return &Router{friendController: *fController, friendService: *fService}
}

//SetupRouter : function that return registered end point
func (r *Router) SetupRouter() *gin.Engine {
	router := gin.New()

	//middleware setup
	router.Use(ginglog.Logger(5), gin.Recovery())

	//diagnostic endpoint
	diagnostic := router.Group("api/v1")
	{
		diagnostic.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Name":       "Friend Management",
				"message":    "OK",
				"serverTime": time.Now().UTC(),
				"version":    "0.0.1",
			})
		})
	}

	//friend endpoint
	friend := router.Group("api/v1/friend")
	{
		friend.POST("/connect", r.friendController.CreateConnection)
	}

	return router
}
