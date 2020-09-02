package domain

// Repository interface for Param struct
type Repository interface {
	Create(model *Param) error
	Update(paramID string, model *Param) error
	GetByID(paramID string) (*Param, error)
	ToggleActive(paramID string) (*Param, error)
	GetAll() (Params, error)
}
