package application

import (
	"database/sql"

	paramAuditDom "github.com/angelmendozacap/go-structure/pkg/paramaudit/domain"
)

// Service for Params
type Service struct {
	Repo paramAuditDom.Repository
}

// NewService creates a new service
func NewService(repo paramAuditDom.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

// CreateTX method creates a new paramAudit
func (s *Service) CreateTX(tx *sql.Tx, paramAudit *paramAuditDom.ParamsAudit) error {
	return s.Repo.CreateTX(tx, paramAudit)
}
