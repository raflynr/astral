package helper

import (
	"github.com/go-playground/validator"
)

func CustomMessageValidator(err error) error {
	var errors []map[string]string

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			field := e.Field()
			tag := e.Tag()

			var msg string
			switch tag {
			case "required":
				msg = field + " is required"
			case "email":
				msg = field + " must be a valid email"
			case "min":
				msg = e.Field() + " must be at least " + e.Param() + " characters"
			case "alphanum":
				msg = field + " must be alphanumeric"
			case "eqfield":
				msg = field + " must be equal to " + e.Param()
			default:
				msg = field + " is not valid"
			}

			errors = append(errors, map[string]string{
				field: msg,
			})
		}

		return NewError("Validation Error", errors)
	}

	return err
}
