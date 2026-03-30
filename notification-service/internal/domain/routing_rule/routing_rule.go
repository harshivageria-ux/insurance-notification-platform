package routing_rule

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"probus-notification-system/internal/domain/validation"
)

// RoutingRule represents a message routing rule
type RoutingRule struct {
	ID                  int       `json:"id" db:"id"`
	TemplateGroupID     int       `json:"template_group_id" db:"template_group_id"`
	ChannelID           int       `json:"channel_id" db:"channel_id"`
	PreferredProviderID int       `json:"preferred_provider_id" db:"preferred_provider_id"`
	FallbackProviderID  int       `json:"fallback_provider_id,omitempty" db:"fallback_provider_id"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	CreatedBy           string    `json:"created_by" db:"created_by"`
	IsActive            bool      `json:"is_active" db:"is_active"`
	Version             int       `json:"version" db:"version"`
}

// CreateRequest represents a request to create a new routing rule
type CreateRequest struct {
	TemplateGroupID     int    `json:"template_group_id"`
	ChannelID           int    `json:"channel_id"`
	PreferredProviderID int    `json:"preferred_provider_id"`
	FallbackProviderID  int    `json:"fallback_provider_id"`
	CreatedBy           string `json:"created_by"`
}

// UpdateRequest represents a request to update a routing rule
type UpdateRequest struct {
	ID                  int    `json:"id"`
	PreferredProviderID int    `json:"preferred_provider_id"`
	FallbackProviderID  int    `json:"fallback_provider_id"`
}

// ToggleRequest represents a request to toggle rule active status
type ToggleRequest struct {
	ID       int  `json:"id"`
	IsActive bool   `json:"is_active"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		TemplateGroupID:     r.TemplateGroupID,
		ChannelID:           r.ChannelID,
		PreferredProviderID: r.PreferredProviderID,
		FallbackProviderID:  r.FallbackProviderID,
		CreatedBy:           validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		ID:                  r.ID,
		PreferredProviderID: r.PreferredProviderID,
		FallbackProviderID:  r.FallbackProviderID,
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.PositiveInt("template group id", r.TemplateGroupID); err != nil {
		return err
	}
	if err := validation.PositiveInt("channel id", r.ChannelID); err != nil {
		return err
	}
	if err := validation.PositiveInt("preferred provider id", r.PreferredProviderID); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.PositiveInt("routing rule id", r.ID); err != nil {
		return err
	}
	return validation.PositiveInt("preferred provider id", r.PreferredProviderID)
}

func normalizeCondition(value json.RawMessage) json.RawMessage {
	trimmed := validation.Trim(string(value))
	if trimmed == "" {
		return nil
	}
	return json.RawMessage(trimmed)
}

// Value implements the driver.Valuer interface
func (r RoutingRule) Value() (driver.Value, error) {
	return json.Marshal(r)
}
