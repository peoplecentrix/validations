package validations

import (
	"fmt"
)

// NewError generate a new error for a model's field
func NewError(field, err string) error {
	return &Error{Field: field, Message: err}
}

// Error is a validation error struct that hold model, field and error message
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error show error message
func (err Error) Error() string {
	return fmt.Sprintf("%v", err.Message)
}
