package main

import (
	"fmt"
	"net/http"

	"github.com/FriendManagement/shared"
	"github.com/FriendManagement/shared/config"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

var (
	runMigration  bool
	db            *gorm.DB
	configuration config.Configuration
	router        *gin.Engine
)

func init() {
	glog.V(2).Info("Initilizing server...")

	//Setup Configuration
	cfg, err := config.New()
	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}
	configuration = *cfg

	//Setup Router
	router = shared.SetupRouter()
}

func main() {
	glog.V(2).Infof("Server run on mode: %s", configuration.Server.Mode)
	gin.SetMode(configuration.Server.Mode)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(fmt.Errorf("Fatal error failed to start server : %s", err))
	}
}
