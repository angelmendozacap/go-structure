package domain

import "time"

// Param model for params
type Param struct {
	ParamID      string    `json:"param_id"`
	Name         string    `json:"name"`
	Value        string    `json:"value"`
	Active       uint8     `json:"active"`
	InsUserID    uint      `json:"ins_user_id"`
	InsDate      time.Time `json:"ins_date"`
	InsDateTime  time.Time `json:"ins_date_time"`
	InsTimestamp int       `json:"ins_timestamp"`
}

// Params Param pointers slice
type Params []*Param
