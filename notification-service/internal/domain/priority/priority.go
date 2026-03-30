package priority

import (
	"time"

	"probus-notification-system/internal/domain/validation"
)

// Priority represents a notification priority level
type Priority struct {
	PriorityID   int16     `json:"priority_id" db:"priority_id"`
	PriorityCode string    `json:"priority_code" db:"priority_code"`
	Description  string    `json:"description" db:"description"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	CreatedBy    string    `json:"created_by" db:"created_by"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	Version      int       `json:"version" db:"version"`
}

// CreateRequest represents a request to create a new priority
type CreateRequest struct {
	PriorityCode string `json:"priority_code"`
	Description  string `json:"description"`
	CreatedBy    string `json:"created_by"`
}

// UpdateRequest represents a request to update a priority
type UpdateRequest struct {
	PriorityID   int16  `json:"priority_id"`
	PriorityCode string `json:"priority_code"`
	Description  string `json:"description"`
	CreatedBy    string `json:"created_by"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		PriorityCode: validation.UpperTrim(r.PriorityCode),
		Description:  validation.Trim(r.Description),
		CreatedBy:    validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		PriorityID:   r.PriorityID,
		PriorityCode: validation.UpperTrim(r.PriorityCode),
		Description:  validation.Trim(r.Description),
		CreatedBy:    validation.Trim(r.CreatedBy),
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.RequireText("priority code", r.PriorityCode, 30); err != nil {
		return err
	}
	if err := validation.OptionalText("priority description", r.Description, 500); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.RequireText("priority code", r.PriorityCode, 30); err != nil {
		return err
	}
	if err := validation.OptionalText("priority description", r.Description, 500); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}
