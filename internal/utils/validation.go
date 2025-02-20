package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
}

func ValidateRequest(input interface{}, customMessages map[string]string) (bool, string) {
	err := Validate.Struct(input)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldErr := range validationErrors {
				fieldName := fieldErr.Field()
				if msg, exists := customMessages[fieldName]; exists {
					errorMessages = append(errorMessages, msg)
				} else {
					// Pesan default jika tidak ada mapping khusus
					errorMessages = append(errorMessages, fieldName+": "+fieldErr.Tag())
				}
			}
			return false, strings.Join(errorMessages, ";")
		}
		return false, err.Error()
	}
	return true, ""
}
