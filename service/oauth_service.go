package service

import (
	"errors"
	"forum-app/entity"
	"forum-app/helper"
	"forum-app/model/request"
	"forum-app/model/response"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"strings"
	"time"
)

type OauthService interface {
	PasswordGrant(request request.PasswordGrant) response.AccessTokenResponse
}

type OauthServiceImpl struct {
	DB *gorm.DB
}

func NewOauthService(DB *gorm.DB) *OauthServiceImpl {
	return &OauthServiceImpl{DB: DB}
}

func (service *OauthServiceImpl) PasswordGrant(request request.PasswordGrant) response.AccessTokenResponse {
	client, err := service.GetClient(request.ClientId)
	helper.PanicIfError(err)

	scope, err := service.GetScope(request.Scope)
	helper.PanicIfError(err)

	user, err := service.AuthUser(request.Username, request.Password)
	helper.PanicIfError(err)

	accessToken, err := service.GetAccessToken(client, user, scope)
	helper.PanicIfError(err)

	refreshToken, err := service.GetRefreshToken(client, user, scope)
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

	var count int64
	service.DB.Model(&entity.OauthScope{}).Where("scope in (?)", scopes).Count(&count)

	if count != int64(len(scopes)) {
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

func (service *OauthServiceImpl) GetAccessToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthAccessToken, error) {
	tx := service.DB.Begin()

	query := tx.Where("client_id = ?", client.ID)
	if user != nil {
		query = query.Where("user_id = ?", user.ID)
	} else {
		query = query.Where("user_id IS NULL")
	}
	if err := query.Where("expires_at <= ?", time.Now()).Delete(new(entity.OauthAccessToken)).Error; err != nil {
		tx.Rollback() // rollback the transaction
		return nil, err
	}

	accessToken := &entity.OauthAccessToken{
		ClientId:  client.ID,
		UserId:    user.ID,
		Token:     uuid.New().String(),
		Scope:     scope,
		ExpiredAt: time.Now().UTC().Add(time.Duration(86400) * time.Second),
	}

	if err := tx.Create(accessToken).Error; err != nil {
		tx.Rollback() // rollback the transaction
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // rollback the transaction
		return nil, err
	}

	return accessToken, nil
}

func (service *OauthServiceImpl) GetRefreshToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthRefreshToken, error) {
	tx := service.DB.Begin()

	refreshToken := &entity.OauthRefreshToken{}
	query := tx.Where("client_id = ?", client.ID)
	if user != nil {
		query = query.Where("user_id = ?", user.ID)
	} else {
		query = query.Where("user_id IS NULL")
	}
	err := query.First(refreshToken).Error
	isFound := !errors.Is(err, gorm.ErrRecordNotFound)
	isNotExpired := false
	if isFound {
		isNotExpired = time.Now().UTC().After(refreshToken.ExpiredAt)
	}

	if isNotExpired {
		return refreshToken, nil
	}

	refreshToken = &entity.OauthRefreshToken{
		ClientId:  client.ID,
		UserId:    user.ID,
		Token:     uuid.New().String(),
		Scope:     scope,
		ExpiredAt: time.Now().UTC().Add(time.Duration(86400) * time.Second),
	}

	if err := tx.Create(refreshToken).Error; err != nil {
		tx.Rollback() // rollback the transaction
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // rollback the transaction
		return nil, err
	}

	return refreshToken, nil
}
