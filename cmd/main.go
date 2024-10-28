package main

import (
	"UnderProject/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.SetupRoutes(router)

	router.Run("8080")
}
