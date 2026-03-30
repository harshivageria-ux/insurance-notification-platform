package language

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"

	"probus-notification-system/internal/domain/validation"
)

// Language represents a language configuration in the system
type Language struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Code      string    `json:"code" db:"code"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedBy string    `json:"updated_by,omitempty" db:"updated_by"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	Version   int       `json:"version" db:"version"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateRequest represents a request to create a new language
type CreateRequest struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    string `json:"status,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
}

// UpdateRequest represents a request to update a language
type UpdateRequest struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    string `json:"status,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

var (
	ErrNameRequired     = errors.New("language name is required")
	ErrNameTooLong      = errors.New("language name must be 100 characters or fewer")
	ErrCodeRequired     = errors.New("language code is required")
	ErrCodeInvalid      = errors.New("language code must contain only letters and be 2 to 10 characters long")
	ErrLanguageNotFound = errors.New("language not found")
	ErrDuplicateName    = errors.New("language name already exists")
	ErrDuplicateCode    = errors.New("language code already exists")
	languageCodePattern = regexp.MustCompile(`^[A-Z]{2,10}$`)
)

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
		Name:      validation.Trim(r.Name),
		Code:      validation.UpperTrim(r.Code),
		Status:    status,
		CreatedBy: createdBy,
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
		ID:        r.ID,
		Name:      validation.Trim(r.Name),
		Code:      validation.UpperTrim(r.Code),
		Status:    status,
		UpdatedBy: updatedBy,
	}
}

func (r CreateRequest) Validate() error {
	if r.Name == "" {
		return ErrNameRequired
	}
	if len([]rune(r.Name)) > 100 {
		return ErrNameTooLong
	}
	if r.Code == "" {
		return ErrCodeRequired
	}
	if !languageCodePattern.MatchString(r.Code) {
		return ErrCodeInvalid
	}
	return nil
}

func (r UpdateRequest) Validate() error {
	if r.ID == 0 {
		return errors.New("language id is required")
	}
	if err := r.Validate(); err != nil {
		return err
	}
	return nil
}

func getIsActiveFromStatus(status string) bool {
	return strings.EqualFold(status, "active")
}


// Value implements the driver.Valuer interface
func (l Language) Value() (driver.Value, error) {
	return json.Marshal(l)
}
