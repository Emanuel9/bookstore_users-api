package app

import (
	"github.com/Emanuel9/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)
func StartApplication() {
	mapURLs()

	logger.Info("about to start the application ...")
	router.Run(":8080")
}