package model

// ServerConfig represents a V2Ray server configuration.
// The fields are structured to hold all necessary data for generating a V2Ray client config.
type ServerConfig struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"` // "vmess", "vless", "shadowsocks", "trojan"
	Address  string `json:"address"`
	Port     int    `json:"port"`

	// VMess/VLESS specific
	UserID   string `json:"userID"`   // The UUID for the user
	AlterID  int    `json:"alterID"`  // Usually 0 for VLESS, >0 for VMess
	Security string `json:"security"` // e.g., "auto", "aes-128-gcm", "none"

	// Transport settings
	Network     string `json:"network"`      // e.g., "tcp", "ws", "grpc"
	WsPath      string `json:"wsPath"`       // Path for WebSocket, e.g., "/path"
	GrpcSvcName string `json:"grpcSvcName"`  // Service name for gRPC

	// TLS settings
	Tls      string `json:"tls"`       // "tls" or ""
	Sni      string `json:"sni"`       // Server Name Indication
}
