package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pomntv/nantavit_agnos_backend/database"
)

func main() {
	database.ConnectDb()

	router := gin.Default()

	setupRoutes(router)

	router.Run(":3006")
}
