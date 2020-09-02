package database

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	// dns cadena de conexión de cada motor de bd
	dnspsql  = "postgres://%s:%s@%s:%d/%s?sslmode=disable"
	dnsmysql = "%s:%s@tcp(%s:%d)/%s?parseTime=true"

	// nombre del motor de bd
	postgresEngine = "postgres"
	mysqlEngine    = "mysql"
)

var (
	once sync.Once
	db   *sql.DB

	// errores
	errNotInitialized = errors.New("el pool de conexiones no se ha inicializado")
)

// Model connection model
type Model struct {
	Engine   string
	User     string
	Password string
	Server   string
	Database string
	Port     int
}

// NewConnection devuelve una única instancia de la conexión
func (m *Model) NewConnection() (*sql.DB, error) {
	var (
		err error
		dns string
	)
	if m.Engine == "" {
		return nil, errors.New("el motor de bases de datos es obligatorio")
	}

	once.Do(func() {
		switch m.Engine {
		case postgresEngine:
			dns = dnspsql
		case mysqlEngine:
			dns = dnsmysql
		}

		db, err = m.getConnection(dns)
	})

	return db, err
}

// getConnection devuelve un pool de conexiones.
func (m *Model) getConnection(dns string) (*sql.DB, error) {
	var err error
	d := fmt.Sprintf(
		dns,
		m.User, m.Password, m.Server, m.Port, m.Database,
	)

	db, err := sql.Open(m.Engine, d)
	if err != nil {
		return db, err
	}

	return db, nil
}

// GetConnection devuelve el pool de conexiones
func GetConnection() (*sql.DB, error) {
	if db == nil {
		return db, errNotInitialized
	}

	if db.Ping() != nil {
		return db, errNotInitialized
	}

	return db, nil
}

// CloseConnection permite cerrar el pool de conexiones
func CloseConnection() error {
	err := db.Close()
	if err != nil {
		return err
	}

	return nil
}
