package request

type PasswordGrant struct {
	ClientId string
	Scope    string
	Username string
	Password string
}
