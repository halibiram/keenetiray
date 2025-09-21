package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"v2ray-keenetic/internal/model"
	"v2ray-keenetic/internal/service"
	"v2ray-keenetic/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTestDB creates a new in-memory SQLite database for testing.
func setupTestDB(t *testing.T) store.Store {
	db, err := store.New(":memory:")
	assert.NoError(t, err)
	return db
}

// newTestServer creates a valid ServerConfig for testing.
func newTestServer(name string) *model.ServerConfig {
	return &model.ServerConfig{
		Name:     name,
		Protocol: "vmess",
		Address:  fmt.Sprintf("%s.com", name),
		Port:     443,
		UserID:   "b8a8a9a8-b8a8-4a8a-8a8a-b8a8a9a8b8a8",
		AlterID:  0,
		Security: "auto",
		Network:  "ws",
		WsPath:   "/",
	}
}

func TestGetServers(t *testing.T) {
	// Setup
	db := setupTestDB(t)
	defer db.Close()

	testServer := newTestServer("TestServer1")
	_, err := db.AddServer(testServer)
	assert.NoError(t, err)

	mockService := service.NewV2RayService()
	handler := NewHandler(mockService, db)

	r := gin.Default()
	r.GET("/api/servers", handler.GetServers)

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/api/servers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var servers []*model.ServerConfig
	err = json.Unmarshal(w.Body.Bytes(), &servers)
	assert.NoError(t, err)

	assert.Len(t, servers, 1)
	assert.Equal(t, "TestServer1", servers[0].Name)
	assert.Equal(t, "vmess", servers[0].Protocol)
}

func TestAddServer(t *testing.T) {
	// Setup
	db := setupTestDB(t)
	defer db.Close()

	mockService := service.NewV2RayService()
	handler := NewHandler(mockService, db)

	r := gin.Default()
	r.POST("/api/servers", handler.AddServer)

	// Prepare request body
	serverToAdd := newTestServer("NewServer")
	body, _ := json.Marshal(serverToAdd)

	// Perform request
	req, _ := http.NewRequest(http.MethodPost, "/api/servers", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var createdServer model.ServerConfig
	err := json.Unmarshal(w.Body.Bytes(), &createdServer)
	assert.NoError(t, err)

	assert.NotEmpty(t, createdServer.ID)
	assert.Equal(t, "NewServer", createdServer.Name)

	// Verify it was actually added to the DB
	servers, err := db.GetAllServers()
	assert.NoError(t, err)
	assert.Len(t, servers, 1)
	assert.Equal(t, "NewServer", servers[0].Name)
}

func TestDeleteServer(t *testing.T) {
	// Setup
	db := setupTestDB(t)
	defer db.Close()

	serverToDelete := newTestServer("ToDelete")
	id, err := db.AddServer(serverToDelete)
	assert.NoError(t, err)

	mockService := service.NewV2RayService()
	handler := NewHandler(mockService, db)

	r := gin.Default()
	r.DELETE("/api/servers/:id", handler.DeleteServer)

	// Perform request
	url := fmt.Sprintf("/api/servers/%d", id)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusNoContent, w.Code)

	servers, err := db.GetAllServers()
	assert.NoError(t, err)
	assert.Len(t, servers, 0)
}
