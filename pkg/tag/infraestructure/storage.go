package infraestructure

import (
	"database/sql"
	"log"

	"github.com/angelmendozacap/go-structure/pkg/tag/application"
	"github.com/angelmendozacap/go-structure/pkg/tag/domain"
	"github.com/angelmendozacap/go-structure/pkg/tag/infraestructure/mysql"
)

const (
	// MySQL is the engine for MySql
	MySQL = "mysql"
)

// NewStore debe invocarse para obtener un ModelStore
// esta función configura el storage para conectarse a la BD
func NewStore(engine string, db *sql.DB) *application.Service {
	serv := &application.Service{}
	serv.Repo = newStorage(engine, db)
	return serv
}

func newStorage(engine string, db *sql.DB) domain.Repository {
	var s domain.Repository

	switch engine {
	case MySQL:
		s = &mysql.Mysql{DB: db}
	default:
		log.Fatalf("el motor de base de datos %s no está implementado.", engine)
	}

	return s
}
