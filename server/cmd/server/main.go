package main

import (
	"database/sql"
	"github.com/G-V-G/l3/server/db"
)

// ServerEnv for port and host
type ServerEnv struct {
	Port int
	Host string
} 

// NewDbConnection gives DB URI
func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "lab3",
		User:       "andrew1407",
		Password:   "sasik",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	senv := &ServerEnv{Port: 5000, Host: "192.168.1.6"}
	server, err := NewServer(senv)
	if err != nil {
		panic(err)
	}
	server.Run()
}
