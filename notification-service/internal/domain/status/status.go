package status

import (
	"errors"
	"strings"
	"time"

	"probus-notification-system/internal/domain/validation"
)

// Status represents a notification status
type Status struct {
	ID          int16     `json:"id" db:"-"`
	StatusID    int16     `json:"status_id" db:"status_id"`
	StatusCode  string    `json:"status_code" db:"status_code"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	IsFinal     bool      `json:"is_final" db:"is_final"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	Version     int       `json:"version" db:"version"`
}

// CreateRequest represents a request to create a new status
type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status,omitempty"`
	IsFinal     bool   `json:"is_final"`
	CreatedBy   string `json:"created_by,omitempty"`
}

// UpdateRequest represents a request to update a status
type UpdateRequest struct {
	StatusID    int16  `json:"status_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status,omitempty"`
	IsFinal     bool   `json:"is_final"`
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
		IsFinal:     r.IsFinal,
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
		StatusID:    r.StatusID,
		Name:        validation.Trim(r.Name),
		Description: validation.Trim(r.Description),
		Status:      status,
		IsFinal:     r.IsFinal,
		UpdatedBy:   updatedBy,
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.RequireText("status name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.OptionalText("status description", r.Description, 500); err != nil {
		return err
	}
	return nil
}

func (r UpdateRequest) Validate() error {
	if r.StatusID <= 0 {
		return errors.New("status_id is required")
	}
	if err := validation.RequireText("status name", r.Name, 100); err != nil {
		return err
	}
	if err := validation.OptionalText("status description", r.Description, 500); err != nil {
		return err
	}
	return nil
}
