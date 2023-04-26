package response

type AccessTokenResponse struct {
	UserID       string `json:"UserID,omitempty"`
	AccessToken  string `json:"accessToken"`
	ExpiresIn    int    `json:"expiresIn"`
	TokenType    string `json:"tokenType"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
