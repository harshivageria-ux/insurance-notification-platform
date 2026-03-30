package provider_setting

import (
	"time"

	"probus-notification-system/internal/domain/validation"
)

// ProviderSetting represents a provider configuration setting
type ProviderSetting struct {
	ID           int       `json:"id" db:"id"`
	ProviderID   int       `json:"provider_id" db:"provider_id"`
	SettingKey   string    `json:"setting_key" db:"setting_key"`
	SettingValue string    `json:"setting_value" db:"setting_value"`
	IsSensitive  bool      `json:"is_sensitive" db:"is_sensitive"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedBy    string    `json:"created_by" db:"created_by"`
	Version      int       `json:"version" db:"version"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// CreateRequest represents a request to create a new provider setting
type CreateRequest struct {
	ProviderID   int    `json:"provider_id"`
	SettingKey   string `json:"setting_key"`
	SettingValue string `json:"setting_value"`
	IsSensitive  bool   `json:"is_sensitive"`
	CreatedBy    string `json:"created_by"`
}

// UpdateRequest represents a request to update a provider setting
type UpdateRequest struct {
	ID           int    `json:"id"`
	SettingKey   string `json:"setting_key"`
	SettingValue string `json:"setting_value"`
	IsSensitive  bool   `json:"is_sensitive"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		ProviderID:   r.ProviderID,
		SettingKey:   validation.Trim(r.SettingKey),
		SettingValue: validation.Trim(r.SettingValue),
		IsSensitive:  r.IsSensitive,
		CreatedBy:    validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		ID:           r.ID,
		SettingKey:   validation.Trim(r.SettingKey),
		SettingValue: validation.Trim(r.SettingValue),
		IsSensitive:  r.IsSensitive,
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.PositiveInt("provider id", r.ProviderID); err != nil {
		return err
	}
	if err := validation.RequireText("setting key", r.SettingKey, 100); err != nil {
		return err
	}
	if err := validation.RequireText("setting value", r.SettingValue, 2000); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.PositiveInt("provider setting id", r.ID); err != nil {
		return err
	}
	if err := validation.RequireText("setting key", r.SettingKey, 100); err != nil {
		return err
	}
	if err := validation.RequireText("setting value", r.SettingValue, 2000); err != nil {
		return err
	}
	return nil
}
