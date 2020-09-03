package domain

import "database/sql"

// Repository interface for Param struct
type Repository interface {
	Create(*Param) error
	UpdateTX(*sql.Tx, string, *Param) error
	GetByID(string) (*Param, error)
	ToggleActive(string) (*Param, error)
	GetAll() (Params, error)
}
