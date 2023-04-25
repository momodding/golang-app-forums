package exception

import "github.com/go-playground/validator/v10"

func ParseTags(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "validateUsername":
		return "Invalid or duplicate username"
	}
	return fe.Error() // default error
}
