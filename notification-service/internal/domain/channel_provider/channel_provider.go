package channel_provider

import (
	"time"
)

// ChannelProvider represents a notification channel provider
type ChannelProvider struct {
	ID           int        `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	ProviderType string     `json:"provider_type" db:"provider_type"`
	IsActive     bool       `json:"is_active" db:"is_active"`
	Status       string     `json:"status" db:"status"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateRequest represents a request to create a new channel provider
type CreateRequest struct {
	Name         string `json:"name"`
	ProviderType string `json:"provider_type"`
	Status       string `json:"status"`
}

// UpdateRequest represents a request to update a channel provider
type UpdateRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProviderType string `json:"provider_type"`
	Status       string `json:"status"`
}

// ToggleRequest represents a request to toggle provider active status
type ToggleRequest struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}
