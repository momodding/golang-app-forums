package request

type AuthorizationGrant struct {
	ClientId  string `json:"client_id"`
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
