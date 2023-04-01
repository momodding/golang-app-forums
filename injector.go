//go:build wireinject
// +build wireinject

package main

import (
	"forum-app/config"
	"forum-app/controller"
	"forum-app/service"
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
