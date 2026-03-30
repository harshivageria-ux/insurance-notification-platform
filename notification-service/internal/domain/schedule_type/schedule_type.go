package schedule_type

import (
	"time"

	"probus-notification-system/internal/domain/validation"
)

// ScheduleType represents a notification schedule type
type ScheduleType struct {
	ScheduleTypeID int16     `json:"schedule_type_id" db:"schedule_type_id"`
	ScheduleCode   string    `json:"schedule_code" db:"schedule_code"`
	Description    string    `json:"description" db:"description"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	Version        int       `json:"version" db:"version"`
}

// CreateRequest represents a request to create a new schedule type
type CreateRequest struct {
	ScheduleCode   string `json:"schedule_code"`
	Description    string `json:"description"`
	CreatedBy      string `json:"created_by"`
}

// UpdateRequest represents a request to update a schedule type
type UpdateRequest struct {
	ScheduleTypeID int16  `json:"schedule_type_id"`
	ScheduleCode   string `json:"schedule_code"`
	Description    string `json:"description"`
}

func (r CreateRequest) Normalize() CreateRequest {
	return CreateRequest{
		ScheduleCode:   validation.UpperTrim(r.ScheduleCode),
		Description:    validation.Trim(r.Description),
		CreatedBy:      validation.Trim(r.CreatedBy),
	}
}

func (r UpdateRequest) Normalize() UpdateRequest {
	return UpdateRequest{
		ScheduleTypeID: r.ScheduleTypeID,
		ScheduleCode:   validation.UpperTrim(r.ScheduleCode),
		Description:    validation.Trim(r.Description),
	}
}

func (r CreateRequest) Validate() error {
	if err := validation.RequireText("schedule code", r.ScheduleCode, 30); err != nil {
		return err
	}
	if err := validation.OptionalText("schedule type description", r.Description, 500); err != nil {
		return err
	}
	return validation.RequireText("created by", r.CreatedBy, 100)
}

func (r UpdateRequest) Validate() error {
	if err := validation.RequireText("schedule code", r.ScheduleCode, 30); err != nil {
		return err
	}
	if err := validation.OptionalText("schedule type description", r.Description, 500); err != nil {
		return err
	}
	return nil
}
