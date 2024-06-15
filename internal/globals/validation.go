package globals

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
)

func validateAlphanumHyphenUnderscore(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, char := range value {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' ||
			char == '_') {
			return false
		}
	}
	return true
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 || len(password) > 255 {
		return false
	}

	var (
		hasUpper bool
		hasLower bool
		hasDigit bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}

func init() {
	Validator = validator.New()
	Validator.RegisterValidation("alphanumerichyphenunderscore", validateAlphanumHyphenUnderscore)
	Validator.RegisterValidation("password", validatePassword)
}
