package service

import (
	"fmt"
	"v2ray-keenetic/internal/model"
)

// V2RayService defines the interface for managing the V2Ray process.
type V2RayService interface {
	Start(config model.ServerConfig) error
	Stop() error
	Restart(config model.ServerConfig) error
	GetStatus() (string, error)
}

// v2rayServiceImpl is a concrete implementation of V2RayService.
type v2rayServiceImpl struct {
	// dependencies like a logger or process manager would go here
}

// NewV2RayService creates a new V2RayService.
func NewV2RayService() V2RayService {
	return &v2rayServiceImpl{}
}

func (s *v2rayServiceImpl) Start(config model.ServerConfig) error {
	fmt.Printf("Mock Start: Starting V2Ray with server: %s\n", config.Name)
	// Real implementation will generate a config and run the v2ray-core binary.
	return nil
}

func (s *v2rayServiceImpl) Stop() error {
	fmt.Println("Mock Stop: Stopping V2Ray...")
	// Real implementation will find and kill the V2Ray process.
	return nil
}

func (s *v2rayServiceImpl) Restart(config model.ServerConfig) error {
	fmt.Println("Mock Restart: Restarting V2Ray...")
	if err := s.Stop(); err != nil {
		return err
	}
	return s.Start(config)
}

func (s *v2rayServiceImpl) GetStatus() (string, error) {
	fmt.Println("Mock GetStatus: Getting V2Ray status...")
	// Real implementation will check if the process is running.
	return "Not Running", nil
}
