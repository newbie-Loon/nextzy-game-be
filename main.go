package main

import (
	"nextzy-game-be/config"
	"nextzy-game-be/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://yourdomain.com", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	r := gin.Default()
	r.Use(cors.New(config))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API running",
		})
	})

	r.POST("/users/create", handlers.CreateUser)
	r.POST("/master/gameCheckpoint", handlers.GetMasterGameCheckpoint)

	r.Run(":8080")
}
