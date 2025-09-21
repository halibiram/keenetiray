package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"v2ray-keenetic/internal/model"
)

// V2RayService defines the interface for managing the V2Ray process.
type V2RayService interface {
	Start(config *model.ServerConfig) error
	Stop() error
	Restart(config *model.ServerConfig) error
	GetStatus() (string, error)
}

// v2rayServiceImpl is a concrete implementation of V2RayService.
type v2rayServiceImpl struct {
	mu         sync.Mutex
	cmd        *exec.Cmd
	configPath string
}

// NewV2RayService creates a new V2RayService.
func NewV2RayService() V2RayService {
	return &v2rayServiceImpl{
		configPath: "/tmp/v2ray_config.json", // Path to store the generated config
	}
}

// Start begins the V2Ray process with a given configuration.
func (s *v2rayServiceImpl) Start(config *model.ServerConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cmd != nil && s.cmd.Process != nil {
		// Check if the process is actually still running
		if err := s.cmd.Process.Signal(syscall.Signal(0)); err == nil {
			return errors.New("V2Ray is already running")
		}
	}

	v2rayConfig, err := s.generateV2RayConfig(config)
	if err != nil {
		return fmt.Errorf("failed to generate v2ray config: %w", err)
	}

	if err := ioutil.WriteFile(s.configPath, v2rayConfig, 0644); err != nil {
		return fmt.Errorf("failed to write v2ray config file: %w", err)
	}

	// Assuming 'v2ray' is in the system's PATH
	s.cmd = exec.Command("v2ray", "-config", s.configPath)
	s.cmd.Stdout = os.Stdout // Pipe stdout to see V2Ray logs
	s.cmd.Stderr = os.Stderr // Pipe stderr for errors

	if err := s.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start v2ray process: %w", err)
	}

	go s.cmd.Wait() // Run in a goroutine to not block

	return nil
}

// Stop terminates the V2Ray process.
func (s *v2rayServiceImpl) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cmd == nil || s.cmd.Process == nil {
		return errors.New("V2Ray is not running")
	}

	if err := s.cmd.Process.Kill(); err != nil {
		// Ignore "process already finished" error
		if err.Error() == "os: process already finished" {
			s.cmd = nil
			return nil
		}
		return fmt.Errorf("failed to kill v2ray process: %w", err)
	}

	s.cmd = nil
	return nil
}

// Restart stops and then starts the V2Ray process.
func (s *v2rayServiceImpl) Restart(config *model.ServerConfig) error {
	// Stop might return an error if not running, which is fine.
	_ = s.Stop()
	return s.Start(config)
}

// GetStatus returns the current status of the V2Ray process.
func (s *v2rayServiceImpl) GetStatus() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cmd != nil && s.cmd.Process != nil {
		// Signal 0 just checks for existence without killing
		if err := s.cmd.Process.Signal(syscall.Signal(0)); err == nil {
			return "Running", nil
		}
	}
	return "Not Running", nil
}

// generateV2RayConfig creates a V2Ray JSON configuration from a ServerConfig model.
func (s *v2rayServiceImpl) generateV2RayConfig(config *model.ServerConfig) ([]byte, error) {
	// Basic inbound for SOCKS5 proxy
	inbound := InboundConfig{
		Tag:      "socks-in",
		Port:     10808, // Local SOCKS5 port
		Listen:   "127.0.0.1",
		Protocol: "socks",
		Settings: &InboundSettings{Auth: "noauth", UDP: true},
	}

	// Outbound settings based on the protocol
	var outboundSettings json.RawMessage
	if config.Protocol == "vmess" {
		settings := VmessOutboundSettings{
			VNext: []VmessServer{
				{
					Address: config.Address,
					Port:    config.Port,
					Users: []VmessUser{
						{
							ID:       config.UserID,
							AlterID:  config.AlterID,
							Security: config.Security,
						},
					},
				},
			},
		}
		settingsBytes, err := json.Marshal(settings)
		if err != nil {
			return nil, err
		}
		outboundSettings = settingsBytes
	} else {
		return nil, fmt.Errorf("protocol %s not yet supported", config.Protocol)
	}

	outbound := OutboundConfig{
		Tag:      "proxy-out",
		Protocol: config.Protocol,
		Settings: &outboundSettings,
		StreamSettings: &StreamSettings{
			Network: config.Network,
			Security: config.Tls,
			WsSettings: &WsSettings{
				Path: config.WsPath,
			},
			GrpcSettings: &GrpcSettings{
				ServiceName: config.GrpcSvcName,
			},
			TlsSettings: &TlsSettings{
				ServerName: config.Sni,
			},
		},
	}

	fullConfig := V2RayConfig{
		Log:       &LogConfig{LogLevel: "warning"},
		Inbounds:  []InboundConfig{inbound},
		Outbounds: []OutboundConfig{outbound},
	}

	return json.MarshalIndent(fullConfig, "", "  ")
}
