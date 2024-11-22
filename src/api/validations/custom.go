package validations

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidationError ساختار خطای اعتبارسنجی
type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

// GetValidationError تابعی برای پردازش خطاهای اعتبارسنجی
func GetValidationError(err error) *[]ValidationError {
	var validationErrors []ValidationError

	// اگر خطا از نوع ValidationErrors باشد
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range ve {

			validationErrors = append(validationErrors, ValidationError{
				Property: e.Field(),
				Tag:      e.Tag(),
				Value:    e.Param(),
				Message:  fmt.Sprintf("Invalid value for %s: %s", e.Field(), e.Tag()),
			})
		}
		return &validationErrors
	}

	return nil
}
