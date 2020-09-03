package domain

import (
	"time"

	"github.com/angelmendozacap/go-structure/pkg/param/domain"
)

// ParamsAudit model for paramsAudit
type ParamsAudit struct {
	AuditID      uint         `json:"auditId"`
	PrevParam    domain.Param `json:"prevParam"`
	Param        domain.Param `json:"param"`
	SetUserID    uint         `json:"setUserId"`
	SetDate      time.Time    `json:"setDate,omitempty"`
	SetDateTime  time.Time    `json:"setDateTime,omitempty"`
	SetTimestamp int64        `json:"setTimestamp,omitempty"`
}
