// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"forum-app/config"
	"forum-app/controller"
	"forum-app/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeCategoryController() *controller.CategoryControllerImpl {
	db := config.NewDbSession()
	categoryServiceImpl := service.NewCategoryService(db)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	return categoryControllerImpl
}

func InitializeUserController() *controller.UserControllerImpl {
	db := config.NewDbSession()
	userServiceImpl := service.NewUserService(db)
	validate := validator.New()
	userControllerImpl := controller.NewUserController(userServiceImpl, validate)
	return userControllerImpl
}

func InitializeOauthController() *controller.OauthControllerImpl {
	db := config.NewDbSession()
	tokenServiceImpl := service.NewTokenService(db)
	oauthServiceImpl := service.NewOauthService(db, tokenServiceImpl)
	validate := validator.New()
	oauthControllerImpl := controller.NewOauthController(oauthServiceImpl, validate)
	return oauthControllerImpl
}

func InitializeAuthController() *controller.AuthControllerImpl {
	db := config.NewDbSession()
	tokenServiceImpl := service.NewTokenService(db)
	oauthServiceImpl := service.NewOauthService(db, tokenServiceImpl)
	authControllerImpl := controller.NewAuthController(oauthServiceImpl)
	return authControllerImpl
}

func InitializeMiddleware() *config.Middleware {
	db := config.NewDbSession()
	tokenServiceImpl := service.NewTokenService(db)
	oauthServiceImpl := service.NewOauthService(db, tokenServiceImpl)
	middleware := config.NewMiddleware(oauthServiceImpl)
	return middleware
}

// injector.go:

var categorySet = wire.NewSet(service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))

var userSet = wire.NewSet(service.NewUserService, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)), controller.NewUserController, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

var oauthSet = wire.NewSet(service.NewOauthService, wire.Bind(new(service.OauthService), new(*service.OauthServiceImpl)), controller.NewOauthController, wire.Bind(new(controller.OauthController), new(*controller.OauthControllerImpl)))

var tokenSet = wire.NewSet(service.NewTokenService, wire.Bind(new(service.TokenService), new(*service.TokenServiceImpl)))

var authSet = wire.NewSet(controller.NewAuthController, wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)))
