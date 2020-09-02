package mysql

import (
	"database/sql"
	"fmt"

	"github.com/angelmendozacap/go-structure/database/mysql"
	"github.com/angelmendozacap/go-structure/pkg/param/domain"
)

// Mysql estructura de conexión a la BD de mysql
type Mysql struct {
	DB *sql.DB
}

var (
	table       = "Params"
	mysqlInsert = fmt.Sprintf(`INSERT INTO %s
		(paramId, name, value, active, insUserId, insDate, insDatetime, insTimestamp)
		VALUES (?, ?, ?, ?, ¨?, ?, ?, ?)`, table)
	mysqlUpdate = fmt.Sprintf("UPDATE %s SET name = ? WHERE paramId = ?", table)
	mysqlGetAll = fmt.Sprintf(`SELECT
		paramId, name, value, active,
		insUserId, insDate, insDatetime, insTimestamp
		FROM %s`, table)
	mysqlGetByID      = mysqlGetAll + " WHERE paramId = ?"
	mysqlUpdateActive = fmt.Sprintf(`UPDATE %s SET
		active = ? WHERE paramId = ?`, table)
)

// Create registra en la BD
func (m *Mysql) Create(param *domain.Param) error {
	stmt, err := m.DB.Prepare(mysqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = mysql.ExecAffectingOneRow(
		stmt,
		param.ParamID,
		param.Name,
		param.Value,
		param.Active,
		param.InsUserID,
		param.InsDate,
		param.InsDateTime,
		param.InsTimestamp,
	)
	if err != nil {
		return err
	}

	return nil
}

// Update actualiza un registro en la BD
func (m *Mysql) Update(paramID string, param *domain.Param) error {
	stmt, err := m.DB.Prepare(mysqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = mysql.ExecAffectingOneRow(
		stmt,
		param.Name,
		paramID,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetByID consulta un registro por su ID
func (m *Mysql) GetByID(paramID string) (*domain.Param, error) {
	stmt, err := m.DB.Prepare(mysqlGetByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return m.scanRow(stmt.QueryRow(paramID))
}

// ToggleActive consulta un registro por su ID
func (m *Mysql) ToggleActive(paramID string) (*domain.Param, error) {
	stmt, err := m.DB.Prepare(mysqlGetByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	param, err := m.scanRow(stmt.QueryRow(paramID))
	if err != nil {
		return nil, err
	}

	if param.Active == 1 {
		param.Active = 0
	} else {
		param.Active = 1
	}

	stmt2, err := m.DB.Prepare(mysqlUpdateActive)
	if err != nil {
		return nil, err
	}
	defer stmt2.Close()

	err = mysql.ExecAffectingOneRow(
		stmt2,
		param.Active,
		param.ParamID,
	)
	if err != nil {
		return nil, err
	}

	return param, nil
}

// GetAll consulta todos los registros de la BD
func (m *Mysql) GetAll() (domain.Params, error) {
	params := make(domain.Params, 0)

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

		params = append(params, model)
	}

	return params, nil
}

func (m *Mysql) scanRow(s mysql.RowScanner) (*domain.Param, error) {
	param := &domain.Param{}

	if err := s.Scan(
		&param.ParamID,
		&param.Name,
		&param.Value,
		&param.Active,
		&param.InsUserID,
		&param.InsDate,
		&param.InsDateTime,
		&param.InsTimestamp,
	); err != nil {
		return param, err
	}

	return param, nil
}
