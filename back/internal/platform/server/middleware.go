package server

import (
	"log"
	"serverpunk/internal/platform/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.Use(gin.Recovery())

	conf := config.GetConfig()

	routing(router)

	setupFrontend(router)

	err := router.Run(":" + conf.Port)
	if err != nil {
		log.Fatalln(err)
	}
}
