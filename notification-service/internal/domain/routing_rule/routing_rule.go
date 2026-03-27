package routing_rule

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// RoutingRule represents a message routing rule
type RoutingRule struct {
	ID            int             `json:"id" db:"id"`
	Name          string          `json:"name" db:"name"`
	Condition     json.RawMessage `json:"condition" db:"condition"`
	TargetChannel int             `json:"target_channel" db:"target_channel"`
	Priority      int             `json:"priority" db:"priority"`
	IsActive      bool            `json:"is_active" db:"is_active"`
	Status        string          `json:"status" db:"status"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateRequest represents a request to create a new routing rule
type CreateRequest struct {
	Name          string          `json:"name"`
	Condition     json.RawMessage `json:"condition"`
	TargetChannel int             `json:"target_channel"`
	Priority      int             `json:"priority"`
	Status        string          `json:"status"`
}

// UpdateRequest represents a request to update a routing rule
type UpdateRequest struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Condition     json.RawMessage `json:"condition"`
	TargetChannel int             `json:"target_channel"`
	Priority      int             `json:"priority"`
	Status        string          `json:"status"`
}

// ToggleRequest represents a request to toggle rule active status
type ToggleRequest struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

// Value implements the driver.Valuer interface
func (r RoutingRule) Value() (driver.Value, error) {
	return json.Marshal(r)
}
