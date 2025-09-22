package main

import (
	"log"
	"net/http"
	"github.com/halibiram/keenetiray/internal/handler"
	"github.com/halibiram/keenetiray/internal/service"
	"github.com/halibiram/keenetiray/internal/store"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-contrib/cors"
	_ "github.com/gorilla/websocket"
	_ "gopkg.in/yaml.v3"
	_ "github.com/fsnotify/fsnotify"
	_ "github.com/shirou/gopsutil/v3"
	_ "github.com/go-playground/validator/v10"
)

func main() {
	// Initialize database
	// The database file will be created in the same directory where the app is run.
	db, err := store.New("v2ray-keenetic.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize services
	v2rayService := service.NewV2RayService()

	// Initialize handlers, passing the database store
	apiHandler := handler.NewHandler(v2rayService, db)

	// Setup router
	r := gin.Default()

	// API routes
	api := r.Group("/api")
	{
		// Server CRUD routes
		serverAPI := api.Group("/servers")
		{
			serverAPI.GET("", apiHandler.GetServers)
			serverAPI.POST("", apiHandler.AddServer)
			serverAPI.PUT("/:id", apiHandler.UpdateServer)
			serverAPI.DELETE("/:id", apiHandler.DeleteServer)
		}

		// Service control routes
		serviceAPI := api.Group("/service")
		{
			serviceAPI.POST("/start", apiHandler.StartService)
			serviceAPI.POST("/stop", apiHandler.StopService)
			serviceAPI.GET("/status", apiHandler.GetServiceStatus)
		}
	}

	// Serve frontend static files. The executable will look for a 'web' directory
	// in the same directory it is run from.
	r.Static("/ui", "./web")
	// Redirect root to the UI
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui")
	})

	// Simple ping route for health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
