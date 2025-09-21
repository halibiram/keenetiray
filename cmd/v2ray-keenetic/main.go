package main

import (
	"v2ray-keenetic/internal/handler"
	"v2ray-keenetic/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize services
	v2rayService := service.NewV2RayService()

	// Initialize handlers
	apiHandler := handler.NewHandler(v2rayService)

	// Setup router
	r := gin.Default()

	// API routes
	api := r.Group("/api")
	{
		api.GET("/servers", apiHandler.GetServers)
		api.POST("/servers", apiHandler.AddServer)
	}

	// Simple ping route for health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	r.Run(":8080")
}
