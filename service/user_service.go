package service

import (
	"errors"
	"forum-app/entity"
	"forum-app/helper"
	"forum-app/model/request"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(request request.UserRegistration) (*entity.OauthUser, error)
}

type UserServiceImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserService(DB *gorm.DB, Validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{DB: DB, Validate: Validate}
}

func (service *UserServiceImpl) Register(request request.UserRegistration) (*entity.OauthUser, error) {
	service.Validate.RegisterValidation("validateUsername", service.ValidateUsername)
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

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
