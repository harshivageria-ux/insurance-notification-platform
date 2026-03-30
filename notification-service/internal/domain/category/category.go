package category

import (
	"errors"
	"strings"
	"time"

	"probus-notification-system/internal/domain/validation"
)

// Category represents a notification category
type Category struct {
	ID          int       `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	Version     int       `json:"version" db:"version"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// CreateRequest represents a request to create a new category
type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
}

// UpdateRequest represents a request to update a category
type UpdateRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status,omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty"`
}

func (r CreateRequest) Normalize() CreateRequest {
	createdBy := validation.Trim(r.CreatedBy)
	if createdBy == "" {
		createdBy = "system"
	}
	status := strings.Title(strings.ToLower(validation.Trim(r.Status)))
	if status != "Inactive" {
		status = "Active"
	}
	return CreateRequest{
		Name:        validation.Trim(r.Name),
		Description: validation.Trim(r.Description),
		Status:      status,
		CreatedBy:   createdBy,
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	updatedBy := validation.Trim(r.UpdatedBy)
	if updatedBy == "" {
		updatedBy = "system"
	}
	status := strings.Title(strings.ToLower(validation.Trim(r.Status)))
	if status != "Inactive" {
		status = "Active"
	}
	return UpdateRequest{
		ID:          r.ID,
		Name:        validation.Trim(r.Name),
		Description: validation.Trim(r.Description),
		Status:      status,
		UpdatedBy:   updatedBy,
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.RequireText("category name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.OptionalText("category description", r.Description, 500); err != nil {
		return err
	}
	return nil
}

func (r UpdateRequest) Validate() error {
	if r.ID <= 0 {
		return errors.New("category id is required")
	}
	if err := validation.RequireText("category name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.OptionalText("category description", r.Description, 500); err != nil {
		return err
	}
	return nil
}
