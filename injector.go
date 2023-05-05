//go:build wireinject
// +build wireinject

package main

import (
	"forum-app/config"
	"forum-app/controller"
	"forum-app/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var categorySet = wire.NewSet(
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializeCategoryController() *controller.CategoryControllerImpl {
	wire.Build(
		config.NewDbSession,
		categorySet,
	)

	return nil
}

var userSet = wire.NewSet(
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserController,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

func InitializeUserController() *controller.UserControllerImpl {
	wire.Build(
		config.NewDbSession,
		validator.New,
		userSet,
	)

	return nil
}

var oauthSet = wire.NewSet(
	service.NewOauthService,
	wire.Bind(new(service.OauthService), new(*service.OauthServiceImpl)),
	controller.NewOauthController,
	wire.Bind(new(controller.OauthController), new(*controller.OauthControllerImpl)),
)

var tokenSet = wire.NewSet(
	service.NewTokenService,
	wire.Bind(new(service.TokenService), new(*service.TokenServiceImpl)),
)

func InitializeOauthController() *controller.OauthControllerImpl {
	wire.Build(
		config.NewDbSession,
		tokenSet,
		validator.New,
		oauthSet,
	)
	return nil
}

var authSet = wire.NewSet(
	controller.NewAuthController,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

func InitializeAuthController() *controller.AuthControllerImpl {
	wire.Build(
		config.NewDbSession,
		tokenSet,
		authSet,
		oauthSet,
	)
	return nil
}

func InitializeMiddleware() *config.Middleware {
	wire.Build(
		config.NewDbSession,
		tokenSet,
		oauthSet,
		config.NewMiddleware,
	)
	return nil
}
