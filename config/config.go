package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/angelmendozacap/go-structure/database"
)

// GetDBInstance returns a new DB instance
func GetDBInstance() (*sql.DB, error) {
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3030"
	}

	port, _ := strconv.Atoi(dbPort)

	conn := &database.Model{
		Database: os.Getenv("DB_DATABASE"),
		Engine:   os.Getenv("DB_CONNECTION"),
		Port:     port,
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USERNAME"),
		Server:   os.Getenv("DB_HOST"),
	}

	return conn.NewConnection()
}

// GetPort returns the app port
func GetPort() string {
	p := os.Getenv("APP_PORT")
	if "" == p {
		p = "3030"
	}

	return fmt.Sprintf(":%s", p)
}
