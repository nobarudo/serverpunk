package server

import (
	"serverpunk/internal/monitor"

	"github.com/gin-gonic/gin"
)

func routing(router *gin.Engine) {

	api := router.Group("/serverpunk/api")
	{
		api.GET("/status", monitor.GetStatus)
	}
}
