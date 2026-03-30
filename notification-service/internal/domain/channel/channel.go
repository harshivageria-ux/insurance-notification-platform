package channel

import (
	"time"

	"probus-notification-system/internal/domain/validation"
)

// Channel represents a notification channel
type Channel struct {
	ID          int       `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	Version     int       `json:"version" db:"version"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// CreateRequest represents a request to create a new channel
type CreateRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

// UpdateRequest represents a request to update a channel
type UpdateRequest struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ToggleRequest represents a request to toggle channel active status
type ToggleRequest struct {
	ID       int  `json:"id"`
	IsActive bool   `json:"is_active"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		Code:        validation.UpperTrim(r.Code),
		Name:        validation.Trim(r.Name),
		Description: validation.Trim(r.Description),
		CreatedBy:   validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		ID:          r.ID,
		Code:        validation.UpperTrim(r.Code),
		Name:        validation.Trim(r.Name),
		Description: validation.Trim(r.Description),
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.RequireText("channel code", r.Code, 30); err != nil {
		return err
	}
	if err := validation.RequireText("channel name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.OptionalText("channel description", r.Description, 500); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.PositiveInt("channel id", r.ID); err != nil {
		return err
	}
	if err := validation.RequireText("channel code", r.Code, 30); err != nil {
		return err
	}
	if err := validation.RequireText("channel name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.OptionalText("channel description", r.Description, 500); err != nil {
		return err
	}
	return nil
}
