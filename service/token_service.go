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
