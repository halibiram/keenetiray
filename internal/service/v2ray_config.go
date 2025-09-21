package service

import "encoding/json"

// This file contains Go structs that map to the V2Ray JSON configuration format.
// See: https://www.v2fly.org/config/

// V2RayConfig is the top-level configuration structure.
type V2RayConfig struct {
	Log      *LogConfig      `json:"log,omitempty"`
	Inbounds []InboundConfig `json:"inbounds"`
	Outbounds []OutboundConfig `json:"outbounds"`
}

// LogConfig represents the log settings.
type LogConfig struct {
	LogLevel string `json:"loglevel"` // "debug", "info", "warning", "error", "none"
}

// InboundConfig represents an inbound connection handler.
type InboundConfig struct {
	Port     int    `json:"port"`
	Listen   string `json:"listen"`
	Protocol string `json:"protocol"`
	Settings *InboundSettings `json:"settings"`
	Tag      string `json:"tag"`
}

// InboundSettings holds settings for an inbound handler.
type InboundSettings struct {
	Auth string `json:"auth,omitempty"`
	UDP  bool   `json:"udp,omitempty"`
	IP   string `json:"ip,omitempty"`
}

// OutboundConfig represents an outbound connection handler.
type OutboundConfig struct {
	Protocol string           `json:"protocol"`
	Settings *json.RawMessage `json:"settings"` // Using RawMessage to handle different protocol settings
	StreamSettings *StreamSettings `json:"streamSettings,omitempty"`
	Tag      string `json:"tag"`
}

// VmessOutboundSettings holds settings for a VMess outbound.
type VmessOutboundSettings struct {
	VNext []VmessServer `json:"vnext"`
}

// VmessServer represents a single VMess server.
type VmessServer struct {
	Address string       `json:"address"`
	Port    int          `json:"port"`
	Users   []VmessUser  `json:"users"`
}

// VmessUser represents a VMess user.
type VmessUser struct {
	ID       string `json:"id"`
	AlterID  int    `json:"alterId"`
	Security string `json:"security"`
}

// StreamSettings configures the transport protocol.
type StreamSettings struct {
	Network   string        `json:"network"`
	Security  string        `json:"security,omitempty"`
	WsSettings *WsSettings  `json:"wsSettings,omitempty"`
	GrpcSettings *GrpcSettings `json:"grpcSettings,omitempty"`
	TlsSettings *TlsSettings `json:"tlsSettings,omitempty"`
}

// WsSettings holds settings for WebSocket transport.
type WsSettings struct {
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

// GrpcSettings holds settings for gRPC transport.
type GrpcSettings struct {
	ServiceName string `json:"serviceName,omitempty"`
}

// TlsSettings holds settings for TLS.
type TlsSettings struct {
	ServerName         string `json:"serverName,omitempty"`
	AllowInsecure      bool   `json:"allowInsecure,omitempty"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify,omitempty"`
}
