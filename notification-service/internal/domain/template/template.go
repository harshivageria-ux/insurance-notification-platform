package template

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"probus-notification-system/internal/domain/validation"
)

// Template represents a notification template
type Template struct {
	ID              int       `json:"id" db:"id"`
	TemplateGroupID int       `json:"template_group_id" db:"template_group_id"`
	ChannelID       int       `json:"channel_id" db:"channel_id"`
	LanguageID      int       `json:"language_id" db:"language_id"`
	TitleTemplate   string    `json:"title_template" db:"title_template"`
	MessageTemplate string    `json:"message_template" db:"message_template"`
	HasVariables    bool      `json:"has_variables" db:"has_variables"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	CreatedBy       string    `json:"created_by" db:"created_by"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	Version         int       `json:"version" db:"version"`
}

// CreateRequest represents a request to create a new template
type CreateRequest struct {
	TemplateGroupID int    `json:"template_group_id"`
	ChannelID       int    `json:"channel_id"`
	LanguageID      int    `json:"language_id"`
	TitleTemplate   string `json:"title_template"`
	MessageTemplate string `json:"message_template"`
	HasVariables    bool   `json:"has_variables"`
	CreatedBy       string `json:"created_by"`
}

// UpdateRequest represents a request to update a template
type UpdateRequest struct {
	ID              int    `json:"id"`
	TemplateGroupID int    `json:"template_group_id"`
	ChannelID       int    `json:"channel_id"`
	LanguageID      int    `json:"language_id"`
	TitleTemplate   string `json:"title_template"`
	MessageTemplate string `json:"message_template"`
	HasVariables    bool   `json:"has_variables"`
}

// PreviewRequest represents a request to preview a template
type PreviewRequest struct {
	ID        int                    `json:"id"`
	Variables map[string]interface{} `json:"variables"`
}

// PreviewResponse represents the preview response
type PreviewResponse struct {
	RenderedContent string `json:"rendered_content"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		TemplateGroupID: r.TemplateGroupID,
		ChannelID:       r.ChannelID,
		LanguageID:      r.LanguageID,
		TitleTemplate:   validation.Trim(r.TitleTemplate),
		MessageTemplate: validation.Trim(r.MessageTemplate),
		HasVariables:    r.HasVariables,
		CreatedBy:       validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		ID:              r.ID,
		TemplateGroupID: r.TemplateGroupID,
		ChannelID:       r.ChannelID,
		LanguageID:      r.LanguageID,
		TitleTemplate:   validation.Trim(r.TitleTemplate),
		MessageTemplate: validation.Trim(r.MessageTemplate),
		HasVariables:    r.HasVariables,
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.PositiveInt("template group id", r.TemplateGroupID); err != nil {
		return err
	}
	if err := validation.PositiveInt("channel id", r.ChannelID); err != nil {
		return err
	}
	if err := validation.PositiveInt("language id", r.LanguageID); err != nil {
		return err
	}
	if err := validation.OptionalText("title template", r.TitleTemplate, 250); err != nil {
		return err
	}
	if err := validation.RequireText("message template", r.MessageTemplate, 5000); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.PositiveInt("template id", r.ID); err != nil {
		return err
	}
	if err := validation.PositiveInt("template group id", r.TemplateGroupID); err != nil {
		return err
	}
	if err := validation.PositiveInt("channel id", r.ChannelID); err != nil {
		return err
	}
	if err := validation.PositiveInt("language id", r.LanguageID); err != nil {
		return err
	}
	if err := validation.OptionalText("title template", r.TitleTemplate, 250); err != nil {
		return err
	}
	return validation.RequireText("message template", r.MessageTemplate, 5000)
}

func normalizeJSON(value json.RawMessage) json.RawMessage {
	trimmed := string(value)
	if validation.Trim(trimmed) == "" {
		return nil
	}
	return json.RawMessage(validation.Trim(trimmed))
}

// Value implements the driver.Valuer interface
func (t Template) Value() (driver.Value, error) {
	return json.Marshal(t)
}
