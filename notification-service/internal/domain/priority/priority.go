package priority

import (
	"time"
)

// Priority represents a notification priority level
type Priority struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Level     int        `json:"level" db:"level"`
	Status    string     `json:"status" db:"status"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateRequest represents a request to create a new priority
type CreateRequest struct {
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Status string `json:"status"`
}

// UpdateRequest represents a request to update a priority
type UpdateRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Status string `json:"status"`
}
