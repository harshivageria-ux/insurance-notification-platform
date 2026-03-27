package template

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Template represents a notification template
type Template struct {
	ID              int             `json:"id" db:"id"`
	TemplateGroupID int             `json:"template_group_id" db:"template_group_id"`
	Name            string          `json:"name" db:"name"`
	Content         string          `json:"content" db:"content"`
	Variables       json.RawMessage `json:"variables" db:"variables"`
	Status          string          `json:"status" db:"status"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CreateRequest represents a request to create a new template
type CreateRequest struct {
	TemplateGroupID int             `json:"template_group_id"`
	Name            string          `json:"name"`
	Content         string          `json:"content"`
	Variables       json.RawMessage `json:"variables"`
	Status          string          `json:"status"`
}

// UpdateRequest represents a request to update a template
type UpdateRequest struct {
	ID              int             `json:"id"`
	TemplateGroupID int             `json:"template_group_id"`
	Name            string          `json:"name"`
	Content         string          `json:"content"`
	Variables       json.RawMessage `json:"variables"`
	Status          string          `json:"status"`
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

// Value implements the driver.Valuer interface
func (t Template) Value() (driver.Value, error) {
	return json.Marshal(t)
}
