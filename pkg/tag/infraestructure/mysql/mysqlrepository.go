package mysql

import (
	"database/sql"
	"fmt"

	"github.com/angelmendozacap/go-structure/database/mysql"
	"github.com/angelmendozacap/go-structure/pkg/tag/domain"
)

// Mysql estructura de conexión a la BD de mysql
type Mysql struct {
	DB *sql.DB
}

var (
	table        = "tags"
	mysqlInsert  = fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", table)
	mysqlUpdate  = fmt.Sprintf("UPDATE %s SET name = ? WHERE id = ?", table)
	mysqlDelete  = fmt.Sprintf("DELETE FROM %s WHERE id = ?", table)
	mysqlGetAll  = fmt.Sprintf("SELECT id, name FROM %s", table)
	mysqlGetByID = mysqlGetAll + " WHERE id = ?"
)

// Create registra en la BD
func (m *Mysql) Create(tag *domain.Tag) error {
	stmt, err := m.DB.Prepare(mysqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		tag.Name,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	tag, err = m.GetByID(uint(id))
	if err != nil {
		return err
	}

	return nil
}

// Update actualiza un registro en la BD
func (m *Mysql) Update(id uint, tag *domain.Tag) error {
	stmt, err := m.DB.Prepare(mysqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = mysql.ExecAffectingOneRow(
		stmt,
		tag.Name,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete elimina un registro de la BD
func (m *Mysql) Delete(id uint) error {
	stmt, err := m.DB.Prepare(mysqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = mysql.ExecAffectingOneRow(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// GetByID consulta un registro por su ID
func (m *Mysql) GetByID(id uint) (*domain.Tag, error) {
	stmt, err := m.DB.Prepare(mysqlGetByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return m.scanRow(stmt.QueryRow(id))
}

// GetAll consulta todos los registros de la BD
func (m *Mysql) GetAll() (domain.Tags, error) {
	tags := make(domain.Tags, 0)

	stmt, err := m.DB.Prepare(mysqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		model, err := m.scanRow(rows)
		if err != nil {
			return nil, err
		}

		tags = append(tags, model)
	}

	return tags, nil
}

func (m *Mysql) scanRow(s mysql.RowScanner) (*domain.Tag, error) {
	tag := &domain.Tag{}

	if err := s.Scan(
		&tag.ID,
		&tag.Name,
	); err != nil {
		return tag, err
	}

	return tag, nil
}
