package request

type AuthorizationGrant struct {
	ClientId  string `json:"clientId"`
	GrantType string `json:"grantType" validate:"required,validateGrantType"`
	Scope     string `json:"scope"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
