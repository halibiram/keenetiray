package handler

import (
	"net/http"
	"strconv"
	"v2ray-keenetic/internal/model"
	"v2ray-keenetic/internal/service"

	"github.com/gin-gonic/gin"
)

// Using a package-level variable for mock data.
// In a real app, this would be handled by a database layer.
var mockServers = []model.ServerConfig{
	{ID: "1", Name: "My First Server", Address: "server1.com", Port: 443},
	{ID: "2", Name: "My Second Server", Address: "server2.com", Port: 8080},
}
var nextID = 3 // Start after the initial mock servers

// Handler holds the dependencies for the API handlers.
type Handler struct {
	V2RayService service.V2RayService
}

// NewHandler creates a new handler with its dependencies.
func NewHandler(v2rayService service.V2RayService) *Handler {
	return &Handler{
		V2RayService: v2rayService,
	}
}

// GetServers returns the list of V2Ray servers.
func (h *Handler) GetServers(c *gin.Context) {
	c.JSON(http.StatusOK, mockServers)
}

// AddServer adds a new V2Ray server.
func (h *Handler) AddServer(c *gin.Context) {
	var newServer model.ServerConfig
	if err := c.ShouldBindJSON(&newServer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// In a real app, generate a unique ID. For this mock, we'll just increment.
	newServer.ID = strconv.Itoa(nextID)
	nextID++
	mockServers = append(mockServers, newServer)
	c.JSON(http.StatusCreated, newServer)
}
