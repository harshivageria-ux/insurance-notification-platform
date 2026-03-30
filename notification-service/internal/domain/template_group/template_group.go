package template_group

import (
	"time"

	"probus-notification-system/internal/domain/validation"
)

// TemplateGroup represents a group of notification templates
type TemplateGroup struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	Description string    `json:"description" db:"description"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	Version     int       `json:"version" db:"version"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// CreateRequest represents a request to create a new template group
type CreateRequest struct {
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

// UpdateRequest represents a request to update a template group
type UpdateRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		Name:        validation.Trim(r.Name),
		CategoryID:  r.CategoryID,
		Description: validation.Trim(r.Description),
		CreatedBy:   validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		ID:          r.ID,
		Name:        validation.Trim(r.Name),
		CategoryID:  r.CategoryID,
		Description: validation.Trim(r.Description),
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.RequireText("template group name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.PositiveInt("category id", r.CategoryID); err != nil {
		return err
	}
	if err := validation.OptionalText("template group description", r.Description, 500); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.PositiveInt("template group id", r.ID); err != nil {
		return err
	}
	if err := validation.RequireText("template group name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.PositiveInt("category id", r.CategoryID); err != nil {
		return err
	}
	if err := validation.OptionalText("template group description", r.Description, 500); err != nil {
		return err
	}
	return nil
}
