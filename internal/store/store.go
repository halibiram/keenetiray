package store

import (
	"database/sql"
	"strconv"
	"v2ray-keenetic/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

// Store defines all the database operations.
type Store interface {
	GetAllServers() ([]*model.ServerConfig, error)
	AddServer(server *model.ServerConfig) (int64, error)
	GetServer(id int64) (*model.ServerConfig, error)
	UpdateServer(server *model.ServerConfig) error
	DeleteServer(id int64) error
	Close() error
}

// DB is the database connection that implements Store.
type DB struct {
	*sql.DB
}

// New creates a new database connection and initializes the schema.
func New(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Updated schema to hold all V2Ray config fields
	query := `
	CREATE TABLE IF NOT EXISTS servers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		protocol TEXT NOT NULL,
		address TEXT NOT NULL,
		port INTEGER NOT NULL,
		user_id TEXT NOT NULL,
		alter_id INTEGER,
		security TEXT,
		network TEXT,
		ws_path TEXT,
		grpc_svc_name TEXT,
		tls TEXT,
		sni TEXT
	);`

	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// GetAllServers retrieves all servers from the database.
func (db *DB) GetAllServers() ([]*model.ServerConfig, error) {
	rows, err := db.Query("SELECT id, name, protocol, address, port, user_id, alter_id, security, network, ws_path, grpc_svc_name, tls, sni FROM servers ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servers []*model.ServerConfig
	for rows.Next() {
		s := &model.ServerConfig{}
		if err := rows.Scan(&s.ID, &s.Name, &s.Protocol, &s.Address, &s.Port, &s.UserID, &s.AlterID, &s.Security, &s.Network, &s.WsPath, &s.GrpcSvcName, &s.Tls, &s.Sni); err != nil {
			return nil, err
		}
		servers = append(servers, s)
	}
	return servers, nil
}

// AddServer adds a new server to the database.
func (db *DB) AddServer(server *model.ServerConfig) (int64, error) {
	stmt := "INSERT INTO servers (name, protocol, address, port, user_id, alter_id, security, network, ws_path, grpc_svc_name, tls, sni) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(stmt, server.Name, server.Protocol, server.Address, server.Port, server.UserID, server.AlterID, server.Security, server.Network, server.WsPath, server.GrpcSvcName, server.Tls, server.Sni)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	server.ID = strconv.FormatInt(id, 10)
	return id, nil
}

// GetServer retrieves a single server by its ID.
func (db *DB) GetServer(id int64) (*model.ServerConfig, error) {
	s := &model.ServerConfig{}
	stmt := "SELECT id, name, protocol, address, port, user_id, alter_id, security, network, ws_path, grpc_svc_name, tls, sni FROM servers WHERE id = ?"
	err := db.QueryRow(stmt, id).Scan(&s.ID, &s.Name, &s.Protocol, &s.Address, &s.Port, &s.UserID, &s.AlterID, &s.Security, &s.Network, &s.WsPath, &s.GrpcSvcName, &s.Tls, &s.Sni)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// UpdateServer updates an existing server in the database.
func (db *DB) UpdateServer(server *model.ServerConfig) error {
	id, err := strconv.ParseInt(server.ID, 10, 64)
	if err != nil {
		return err
	}
	stmt := "UPDATE servers SET name=?, protocol=?, address=?, port=?, user_id=?, alter_id=?, security=?, network=?, ws_path=?, grpc_svc_name=?, tls=?, sni=? WHERE id=?"
	_, err = db.Exec(stmt, server.Name, server.Protocol, server.Address, server.Port, server.UserID, server.AlterID, server.Security, server.Network, server.WsPath, server.GrpcSvcName, server.Tls, server.Sni, id)
	return err
}

// DeleteServer removes a server from the database.
func (db *DB) DeleteServer(id int64) error {
	_, err := db.Exec("DELETE FROM servers WHERE id = ?", id)
	return err
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.DB.Close()
}
