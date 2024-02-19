package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pomntv/nantavit_agnos_backend/handlers"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/", handlers.ListPassword)

	router.POST("/password", handlers.CreatePassword)
}
