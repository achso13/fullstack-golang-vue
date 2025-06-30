package main

import (
	"backend/config"
	"backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	database.InitDB()

	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
