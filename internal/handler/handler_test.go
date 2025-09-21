package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"v2ray-keenetic/internal/model"
	"v2ray-keenetic/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouterForTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestGetServers(t *testing.T) {
	// Setup
	// We use the mock service which doesn't have any dependencies for now.
	mockService := service.NewV2RayService()
	handler := NewHandler(mockService)

	r := setupRouterForTest()
	r.GET("/api/servers", handler.GetServers)

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/api/servers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var servers []model.ServerConfig
	err := json.Unmarshal(w.Body.Bytes(), &servers)
	assert.NoError(t, err)

	// Check if the response contains the mock data defined in handler.go
	assert.Len(t, servers, 2)
	assert.Equal(t, "My First Server", servers[0].Name)
	assert.Equal(t, "server1.com", servers[0].Address)
}
