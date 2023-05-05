package service

import (
	"errors"
	"forum-app/entity"
	"forum-app/helper"
	"forum-app/model/request"
	"forum-app/model/response"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"strings"
)

type OauthService interface {
	ValidateGrantType(field validator.FieldLevel) bool
	AuthorizeCodeGrant(request request.AuthorizationGrant) response.AccessTokenResponse
	PasswordGrant(request request.AuthorizationGrant) response.AccessTokenResponse
	RefreshTokenGrant(request request.AuthorizationGrant) response.AccessTokenResponse
	GetClient(clientId string) (*entity.OauthClient, error)
	GetScope(requestScope string) (string, error)
	AuthUser(username string, password string) (*entity.OauthUser, error)
	GetToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthAccessToken, *entity.OauthRefreshToken, error)
}

type OauthServiceImpl struct {
	DB           *gorm.DB
	tokenService TokenService
}

func NewOauthService(DB *gorm.DB, tokenService TokenService) *OauthServiceImpl {
	return &OauthServiceImpl{DB: DB, tokenService: tokenService}
}

func (service *OauthServiceImpl) AuthorizeCodeGrant(request request.AuthorizationGrant) response.AccessTokenResponse {
	return response.AccessTokenResponse{}
}

func (service *OauthServiceImpl) PasswordGrant(request request.AuthorizationGrant) response.AccessTokenResponse {
	client, err := service.GetClient(request.ClientId)
	helper.PanicIfError(err)

	scope, err := service.GetScope(request.Scope)
	helper.PanicIfError(err)

	user, err := service.AuthUser(request.Username, request.Password)
	helper.PanicIfError(err)

	accessToken, err := service.tokenService.GetAccessToken(client, user, scope)
	helper.PanicIfError(err)

	refreshToken, err := service.tokenService.GetRefreshToken(client, user, scope)
	helper.PanicIfError(err)

	return response.AccessTokenResponse{
		UserID:       strconv.FormatUint(user.ID, 10),
		AccessToken:  accessToken.Token,
		ExpiresIn:    86400,
		TokenType:    "Bearer",
		Scope:        scope,
		RefreshToken: refreshToken.Token,
	}
}

func (service *OauthServiceImpl) RefreshTokenGrant(request request.AuthorizationGrant) response.AccessTokenResponse {
	client, err := service.GetClient(request.ClientId)
	helper.PanicIfError(err)

	refreshToken, err := service.tokenService.GetRefreshTokenByToken(request.RefreshToken, client)
	helper.PanicIfError(err)

	scope, err := service.GetScope(request.Scope)
	helper.PanicIfError(err)

	user := &entity.OauthUser{}
	err = service.DB.Where("id = ?", refreshToken.UserId).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		helper.PanicIfError(errors.New("user not found"))
	}

	accessToken, err := service.tokenService.GetAccessToken(client, user, scope)
	helper.PanicIfError(err)

	return response.AccessTokenResponse{
		UserID:       strconv.FormatUint(user.ID, 10),
		AccessToken:  accessToken.Token,
		ExpiresIn:    86400,
		TokenType:    "Bearer",
		Scope:        scope,
		RefreshToken: refreshToken.Token,
	}
}

func (service *OauthServiceImpl) GetClient(clientId string) (*entity.OauthClient, error) {
	client := &entity.OauthClient{}
	err := service.DB.Where("key = LOWER(?)", clientId).First(client).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("client not found")
	}

	return client, nil
}

func (service *OauthServiceImpl) GetScope(requestScope string) (string, error) {
	if requestScope == "" {
		var scopes []string
		service.DB.Model(&entity.OauthScope{}).Where("is_default = ?", true).Pluck("scope", &scopes)

		sort.Strings(scopes)
		return strings.Join(scopes, " "), nil
	}

	scopes := strings.Split(requestScope, " ")
	countOfScopes := int64(len(scopes))

	var count int64
	service.DB.Model(&entity.OauthScope{}).Where("scope in (?)", scopes).Count(&count)

	if count == countOfScopes {
		return requestScope, nil
	}

	return "", errors.New("invalid Scope")
}

func (service *OauthServiceImpl) AuthUser(username string, password string) (*entity.OauthUser, error) {
	user := &entity.OauthUser{}
	result := service.DB.Where("username = LOWER(?)", username).First(user).Error
	if result != nil && errors.Is(result, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	isValidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if isValidPassword != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

func (service *OauthServiceImpl) GetToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthAccessToken, *entity.OauthRefreshToken, error) {
	accessToken, err := service.tokenService.GetAccessToken(client, user, scope)
	helper.PanicIfError(err)

	refreshToken, err := service.tokenService.GetRefreshToken(client, user, scope)
	helper.PanicIfError(err)

	return accessToken, refreshToken, nil
}

func (service *OauthServiceImpl) ValidateGrantType(field validator.FieldLevel) bool {
	grantType := map[string]int{
		"password":      1,
		"refresh_token": 1,
	}

	_, isGrantTypeExist := grantType[field.Field().String()]
	return isGrantTypeExist
}
