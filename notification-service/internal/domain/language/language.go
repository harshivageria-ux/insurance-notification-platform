package language

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Language represents a language configuration in the system
type Language struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Code      string     `json:"code" db:"code"`
	Status    string     `json:"status" db:"status"` // Active, Inactive
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateRequest represents a request to create a new language
type CreateRequest struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

// UpdateRequest represents a request to update a language
type UpdateRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

// Value implements the driver.Valuer interface
func (l Language) Value() (driver.Value, error) {
	return json.Marshal(l)
}
