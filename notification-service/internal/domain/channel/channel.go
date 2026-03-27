package channel

import (
	"time"
)

// Channel represents a notification channel
type Channel struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Type      string     `json:"channel_type" db:"channel_type"`
	IsActive  bool       `json:"is_active" db:"is_active"`
	Status    string     `json:"status" db:"status"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateRequest represents a request to create a new channel
type CreateRequest struct {
	Name   string `json:"name"`
	Type   string `json:"channel_type"`
	Status string `json:"status"`
}

// UpdateRequest represents a request to update a channel
type UpdateRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"channel_type"`
	Status string `json:"status"`
}

// ToggleRequest represents a request to toggle channel active status
type ToggleRequest struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}
