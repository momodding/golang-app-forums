package service

import (
	"errors"
	"forum-app/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TokenService interface {
	GetAccessToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthAccessToken, error)
	GetRefreshToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthRefreshToken, error)
	GetRefreshTokenByToken(token string, client *entity.OauthClient) (*entity.OauthRefreshToken, error)
	GetAccessTokenByToken(token string) (*entity.OauthAccessToken, error)
}

type TokenServiceImpl struct {
	DB *gorm.DB
}

func NewTokenService(DB *gorm.DB) *TokenServiceImpl {
	return &TokenServiceImpl{DB: DB}
}

func (service *TokenServiceImpl) GetAccessToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthAccessToken, error) {
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

func (service *TokenServiceImpl) GetRefreshTokenByToken(token string, client *entity.OauthClient) (*entity.OauthRefreshToken, error) {
	refreshToken := &entity.OauthRefreshToken{}
	err := service.DB.Where("client_id = ?", client.ID).
		Where("token = ?", token).First(refreshToken).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("refresh token not found")
	}

	if time.Now().UTC().After(refreshToken.ExpiredAt) {
		return nil, errors.New("refresh token expired")
	}

	return refreshToken, nil
}

func (service *TokenServiceImpl) GetRefreshToken(client *entity.OauthClient, user *entity.OauthUser, scope string) (*entity.OauthRefreshToken, error) {
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
	isExpired := false
	if isFound {
		isExpired = time.Now().UTC().After(refreshToken.ExpiredAt)
	}

	if !isExpired {
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

func (service *TokenServiceImpl) GetAccessTokenByToken(token string) (*entity.OauthAccessToken, error) {
	accessToken := &entity.OauthAccessToken{}
	err := service.DB.Where("token = ?", token).First(accessToken).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if time.Now().UTC().After(accessToken.ExpiredAt) {
		return nil, errors.New("access token expired")
	}

	return accessToken, nil
}
