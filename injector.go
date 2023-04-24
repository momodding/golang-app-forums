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
