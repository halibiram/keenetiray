package model

// ServerConfig represents a V2Ray server configuration.
// It will be expanded with more fields from V2Ray spec.
type ServerConfig struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}
