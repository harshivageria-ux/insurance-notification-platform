package channel_provider

import (
	"time"

	"probus-notification-system/internal/domain/validation"
)

// ChannelProvider represents a notification channel provider
type ChannelProvider struct {
	ID        int       `json:"id" db:"id"`
	ChannelID int       `json:"channel_id" db:"channel_id"`
	Name      string    `json:"name" db:"name"`
	Code      string    `json:"code" db:"code"`
	Priority  int       `json:"priority" db:"priority"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	Version   int       `json:"version" db:"version"`
}

// CreateRequest represents a request to create a new channel provider
type CreateRequest struct {
	ChannelID int    `json:"channel_id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Priority  int    `json:"priority"`
	CreatedBy string `json:"created_by"`
}

// UpdateRequest represents a request to update a channel provider
type UpdateRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// ToggleRequest represents a request to toggle provider active status
type ToggleRequest struct {
	ID       int  `json:"id"`
	IsActive bool   `json:"is_active"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		ChannelID: r.ChannelID,
		Name:      validation.Trim(r.Name),
		Code:      validation.UpperTrim(r.Code),
		Priority:  r.Priority,
		CreatedBy: validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		ID:       r.ID,
		Name:     validation.Trim(r.Name),
		Priority: r.Priority,
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.PositiveInt("channel id", r.ChannelID); err != nil {
		return err
	}
	if err := validation.RequireText("channel provider name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.RequireText("channel provider code", r.Code, 30); err != nil {
		return err
	}
	if err := validation.PositiveInt("provider priority", r.Priority); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.PositiveInt("channel provider id", r.ID); err != nil {
		return err
	}
	if err := validation.RequireText("channel provider name", r.Name, 100); err != nil {
		return err
	}
	return validation.PositiveInt("provider priority", r.Priority)
}
