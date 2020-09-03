package domain

import (
	"database/sql"
)

// Repository interface for ParamAudit struct
type Repository interface {
	CreateTX(*sql.Tx, *ParamsAudit) error
}
