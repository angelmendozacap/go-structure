package mysql

import (
	"database/sql"
	"fmt"

	"github.com/angelmendozacap/go-structure/database/mysql"
	paramAuditDom "github.com/angelmendozacap/go-structure/pkg/paramaudit/domain"
	"github.com/angelmendozacap/go-structure/utils"
)

// Mysql estructura de conexión a la BD de mysql
type Mysql struct {
	DB *sql.DB
}

var (
	table       = "ParamsAudit"
	mysqlInsert = fmt.Sprintf(`
		INSERT INTO %s (
			paramId, namePrev, valuePrev, activePrev, name, value, active, setUserId,
			setDate, setDatetime, setTimestamp
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, table)
)

// CreateTX registra en la BD
func (m *Mysql) CreateTX(tx *sql.Tx, paramAudit *paramAuditDom.ParamsAudit) error {
	stmt, err := tx.Prepare(mysqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := utils.Now()
	err = mysql.ExecAffectingOneRow(
		stmt,
		paramAudit.PrevParam.ParamID,
		paramAudit.PrevParam.Name,
		paramAudit.PrevParam.Value,
		paramAudit.PrevParam.Active,
		paramAudit.Param.Name,
		paramAudit.Param.Value,
		paramAudit.Param.Active,
		paramAudit.SetUserID,
		now["date"], now["time"], now["unix"],
	)
	if err != nil {
		return err
	}

	return nil
}
