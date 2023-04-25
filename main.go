package main

import (
	"forum-app/config"
	"forum-app/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
	app.Use(recover.New())

	config.LoadConfig()

	categoryController := InitializeCategoryController()
	userController := InitializeUserController()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/categories", categoryController.FindAll)
	app.Post("/categories", categoryController.Create)

	app.Post("/users/register", userController.Register)

	app.Listen(":" + viper.GetString("app.port"))
}
