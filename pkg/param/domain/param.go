package domain

import "time"

// Param model for params
type Param struct {
	ParamID      string    `json:"paramId"`
	Name         string    `json:"name"`
	Value        string    `json:"value"`
	Active       uint8     `json:"active"`
	InsUserID    uint      `json:"insUserId"`
	InsDate      time.Time `json:"insDate,omitempty"`
	InsDateTime  time.Time `json:"insDateTime,omitempty"`
	InsTimestamp int64     `json:"insTimestamp,omitempty"`
}

// Params Param pointers slice
type Params []*Param
