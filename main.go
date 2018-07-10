package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	runMigration bool
	db           *gorm.DB
	router       *gin.Engine
)

func init() {

}

func main() {
	gin.SetMode("Debug")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(fmt.Errorf("Fatal error failed to start server : %s", err))
	}
}
