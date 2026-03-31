package domain

import "time"

// NotificationCategoryChannel represents mapping between category and channel
type NotificationCategoryChannel struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	ChannelID  int       `json:"channel_id"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ChannelProviderMasterMap represents mapping between channel and provider
type ChannelProviderMasterMap struct {
	ID         int       `json:"id"`
	ChannelID  int       `json:"channel_id"`
	ProviderID int       `json:"provider_id"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TemplateChannelLanguageMasterMap represents template->channel->language mapping
type TemplateChannelLanguageMasterMap struct {
	ID         int       `json:"id"`
	TemplateID int       `json:"template_id"`
	ChannelID  int       `json:"channel_id"`
	LanguageID int       `json:"language_id"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Create requests

type CreateNotificationCategoryChannelRequest struct {
	CategoryID int `json:"category_id" binding:"required"`
	ChannelID  int `json:"channel_id" binding:"required"`
}

type CreateChannelProviderMasterMapRequest struct {
	ChannelID  int `json:"channel_id" binding:"required"`
	ProviderID int `json:"provider_id" binding:"required"`
}

type CreateTemplateChannelLanguageMasterMapRequest struct {
	TemplateID int `json:"template_id" binding:"required"`
	ChannelID  int `json:"channel_id" binding:"required"`
	LanguageID int `json:"language_id" binding:"required"`
}
