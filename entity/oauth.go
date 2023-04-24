package entity

import "time"

type OauthClient struct {
	CommonEntity `gorm:"embedded"`
	Key          string `gorm:"column:key;size:255" json:"key"`
	Secret       string `gorm:"column:secret;size:100" json:"secret"`
	RedirectUri  string `gorm:"column:redirect_uri;size:255" json:"redirect_uri"`
}

func (table *OauthClient) TableName() string {
	return "oauth_client"
}

type OauthScope struct {
	CommonEntity `gorm:"embedded"`
	Scope        string `gorm:"column:scope;size:255" json:"scope"`
	Description  string `gorm:"column:description;size:100" json:"Description"`
	IsDefault    bool   `gorm:"column:is_default;size:255;default:false"`
}

func (table *OauthScope) TableName() string {
	return "oauth_scope"
}

type OauthRole struct {
	CommonEntity `gorm:"embedded"`
	Name         string `gorm:"column:name;size:50" json:"name"`
}

func (table *OauthRole) TableName() string {
	return "oauth_role"
}

type OauthUser struct {
	CommonEntity `gorm:"embedded"`
	RoleId       uint64 `gorm:"column:role_id" json:"role_id"`
	Username     string `gorm:"column:username;size:50" json:"username"`
	Password     string `gorm:"column:password;size:20" json:"password"`
}

func (table *OauthUser) TableName() string {
	return "oauth_user"
}

type OauthRefreshToken struct {
	CommonEntity `gorm:"embedded"`
	ClientId     uint64    `gorm:"column:client_id" json:"client_id"`
	UserId       uint64    `gorm:"column:user_id" json:"user_id"`
	Token        string    `gorm:"column:token;size:255" json:"token"`
	Scope        string    `gorm:"column:scope;size:100" json:"scope"`
	ExpiredAt    time.Time `gorm:"column:expires_at" json:"expire_at"`
}

func (table *OauthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}

type OauthAccessToken struct {
	CommonEntity `gorm:"embedded"`
	ClientId     uint64    `gorm:"column:client_id" json:"client_id"`
	UserId       uint64    `gorm:"column:user_id" json:"user_id"`
	Token        string    `gorm:"column:token;size:255" json:"token"`
	Scope        string    `gorm:"column:scope;size:100" json:"scope"`
	ExpiredAt    time.Time `gorm:"column:expires_at" json:"expire_at"`
}

func (table *OauthAccessToken) TableName() string {
	return "oauth_access_token"
}

type OauthAuthCode struct {
	CommonEntity `gorm:"embedded"`
	ClientId     uint64 `gorm:"column:client_id" json:"client_id"`
	UserId       uint64 `gorm:"column:user_id" json:"user_id"`
	Code         string `gorm:"column:code;size:255" json:"code"`
	RedirectUri  string `gorm:"column:redirect_uri;size:255" json:"redirect_uri"`
	Scope        string `gorm:"column:scope;size:100" json:"scope"`
}

func (table *OauthAuthCode) TableName() string {
	return "oauth_auth_token"
}
