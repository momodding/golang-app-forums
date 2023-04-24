package request

type UserRegistration struct {
	Username string `json:"username" validate:"required,validateUsername"`
	Password string `json:"password" validate:"required"`
}
