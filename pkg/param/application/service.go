package application

import "github.com/angelmendozacap/go-structure/pkg/param/domain"

// Service for Params
type Service struct {
	Repo domain.Repository
}

// NewService creates a new service
func NewService(repo domain.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

// Create method creates a new param
func (s *Service) Create(m *domain.Param) error {
	return s.Repo.Create(m)
}

// Update method updates a param
func (s *Service) Update(paramID string, m *domain.Param) error {
	return s.Repo.Update(paramID, m)
}

// GetByID method retrieves a param
func (s *Service) GetByID(paramID string) (*domain.Param, error) {
	return s.Repo.GetByID(paramID)
}

// ToggleActive method retrieves a param and toggle active
func (s *Service) ToggleActive(paramID string) (*domain.Param, error) {
	return s.Repo.ToggleActive(paramID)
}

// GetAll method retrieves all the params
func (s *Service) GetAll() (domain.Params, error) {
	return s.Repo.GetAll()
}
