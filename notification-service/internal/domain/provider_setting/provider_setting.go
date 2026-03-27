package provider_setting

import (
	"time"
)

// ProviderSetting represents a provider configuration setting
type ProviderSetting struct {
	ID           int        `json:"id" db:"id"`
	ProviderID   int        `json:"provider_id" db:"provider_id"`
	SettingKey   string     `json:"setting_key" db:"setting_key"`
	SettingValue string     `json:"setting_value" db:"setting_value"`
	Description  string     `json:"description" db:"description"`
	Status       string     `json:"status" db:"status"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateRequest represents a request to create a new provider setting
type CreateRequest struct {
	ProviderID   int    `json:"provider_id"`
	SettingKey   string `json:"setting_key"`
	SettingValue string `json:"setting_value"`
	Description  string `json:"description"`
	Status       string `json:"status"`
}

// UpdateRequest represents a request to update a provider setting
type UpdateRequest struct {
	ID           int    `json:"id"`
	ProviderID   int    `json:"provider_id"`
	SettingKey   string `json:"setting_key"`
	SettingValue string `json:"setting_value"`
	Description  string `json:"description"`
	Status       string `json:"status"`
}
