package validation

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Trim(value string) string {
	return strings.TrimSpace(value)
}

func UpperTrim(value string) string {
	return strings.ToUpper(Trim(value))
}

func RequireText(fieldLabel, value string, maxLength int) error {
	trimmed := Trim(value)
	if trimmed == "" {
		return fmt.Errorf("%s is required", fieldLabel)
	}

	if maxLength > 0 && len([]rune(trimmed)) > maxLength {
		return fmt.Errorf("%s must be %d characters or fewer", fieldLabel, maxLength)
	}

	return nil
}

func OptionalText(fieldLabel, value string, maxLength int) error {
	trimmed := Trim(value)
	if trimmed == "" {
		return nil
	}

	if maxLength > 0 && len([]rune(trimmed)) > maxLength {
		return fmt.Errorf("%s must be %d characters or fewer", fieldLabel, maxLength)
	}

	return nil
}

func PositiveInt(fieldLabel string, value int) error {
	if value <= 0 {
		return fmt.Errorf("%s must be greater than 0", fieldLabel)
	}
	return nil
}

func NonNegativeInt(fieldLabel string, value int) error {
	if value < 0 {
		return fmt.Errorf("%s cannot be negative", fieldLabel)
	}
	return nil
}

func Status(value string) error {
	trimmed := Trim(value)
	if trimmed != "Active" && trimmed != "Inactive" {
		return fmt.Errorf("status must be either Active or Inactive")
	}
	return nil
}

func JSON(fieldLabel string, value json.RawMessage, required bool) error {
	trimmed := strings.TrimSpace(string(value))
	if trimmed == "" {
		if required {
			return fmt.Errorf("%s is required", fieldLabel)
		}
		return nil
	}

	var decoded any
	if err := json.Unmarshal([]byte(trimmed), &decoded); err != nil {
		return fmt.Errorf("%s must be valid JSON", fieldLabel)
	}

	return nil
}
