package domain

type Repository interface {
	Create(model *Tag) error
	Update(ID uint, model *Tag) error
	Delete(ID uint) error
	GetByID(ID uint) (*Tag, error)
	GetAll() (Tags, error)
}
