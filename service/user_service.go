package service

import (
	"errors"
	"forum-app/entity"
	"forum-app/model/request"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(request request.UserRegistration) (*entity.OauthUser, error)
	ValidateUsername(field validator.FieldLevel) bool
}

type UserServiceImpl struct {
	DB *gorm.DB
}

func NewUserService(DB *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{DB: DB}
}

func (service *UserServiceImpl) Register(request request.UserRegistration) (*entity.OauthUser, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 3)
	if err != nil {
		return nil, err
	}

	user := &entity.OauthUser{
		Username: request.Username,
		Password: string(password),
		RoleId:   1,
	}

	if err := service.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserServiceImpl) FindUserByUsername(username string) (*entity.OauthUser, error) {
	user := &entity.OauthUser{}
	err := service.DB.Where("username = LOWER(?)", username).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (service *UserServiceImpl) ValidateUsername(field validator.FieldLevel) bool {
	_, err := service.FindUserByUsername(field.Field().String())
	if err != nil {
		return false
	}

	return true
}
