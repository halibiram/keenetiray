package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"v2ray-keenetic/internal/model"
	"v2ray-keenetic/internal/service"
	"v2ray-keenetic/internal/store"

	"github.com/gin-gonic/gin"
)

// Handler holds the dependencies for the API handlers.
type Handler struct {
	V2RayService service.V2RayService
	Store        store.Store
}

// NewHandler creates a new handler with its dependencies.
func NewHandler(v2rayService service.V2RayService, store store.Store) *Handler {
	return &Handler{
		V2RayService: v2rayService,
		Store:        store,
	}
}

// GetServers returns the list of V2Ray servers from the database.
func (h *Handler) GetServers(c *gin.Context) {
	servers, err := h.Store.GetAllServers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve servers"})
		return
	}
	// If servers is nil (no rows), return an empty array instead of null
	if servers == nil {
		servers = make([]*model.ServerConfig, 0)
	}
	c.JSON(http.StatusOK, servers)
}

// AddServer adds a new V2Ray server to the database.
func (h *Handler) AddServer(c *gin.Context) {
	var newServer model.ServerConfig
	if err := c.ShouldBindJSON(&newServer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.Store.AddServer(&newServer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add server"})
		return
	}

	c.JSON(http.StatusCreated, newServer)
}

// DeleteServer removes a server from the database.
func (h *Handler) DeleteServer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server ID"})
		return
	}

	if err := h.Store.DeleteServer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete server"})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateServer updates an existing server.
func (h *Handler) UpdateServer(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server ID"})
		return
	}

	var serverToUpdate model.ServerConfig
	if err := c.ShouldBindJSON(&serverToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the ID from the URL is assigned to the model
	serverToUpdate.ID = idStr

	if err := h.Store.UpdateServer(&serverToUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update server"})
		return
	}

	c.JSON(http.StatusOK, serverToUpdate)
}

// StartService starts the V2Ray service with a given server config.
func (h *Handler) StartService(c *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, missing server ID"})
		return
	}

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server ID format"})
		return
	}

	config, err := h.Store.GetServer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server configuration not found"})
		return
	}

	if err := h.V2RayService.Start(config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to start V2Ray service: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "V2Ray service started successfully"})
}

// StopService stops the V2Ray service.
func (h *Handler) StopService(c *gin.Context) {
	if err := h.V2RayService.Stop(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to stop V2Ray service: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "V2Ray service stopped successfully"})
}

// GetServiceStatus returns the current status of the V2Ray service.
func (h *Handler) GetServiceStatus(c *gin.Context) {
	status, err := h.V2RayService.GetStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get V2Ray service status: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}
