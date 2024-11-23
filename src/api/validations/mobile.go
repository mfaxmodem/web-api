package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/mfaxmodem/web-api/src/common"
)

func MobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok || value == "" {
		return false
	}
	return common.MobileNumberValidator(value)
}
