package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	fmt.Sprintln("hey ho")
	return strings.Contains(field.Field().String(), "cool")
}