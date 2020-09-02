package application

import "github.com/angelmendozacap/go-structure/pkg/tag/domain"

type Service struct {
	Repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

// Create método para crear un registro de user
func (s *Service) Create(m *domain.Tag) error {
	return s.Repo.Create(m)
}

// Update método para actualizar un registro de user
func (s *Service) Update(ID uint, m *domain.Tag) error {
	return s.Repo.Update(ID, m)
}

// Delete método para eliminar un registro de user
func (s *Service) Delete(ID uint) error {
	return s.Repo.Delete(ID)
}

// GetByID método para obtener un registro de user
func (s *Service) GetByID(ID uint) (*domain.Tag, error) {
	return s.Repo.GetByID(ID)
}

// GetAll método para obtener todos los registro de user
func (s *Service) GetAll() (domain.Tags, error) {
	return s.Repo.GetAll()
}
