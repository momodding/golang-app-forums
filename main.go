package main

import (
	"forum-app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()

	config.LoadConfig()

	categoryController := InitializeCategoryController()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/categories", categoryController.FindAll)
	app.Post("/categories", categoryController.Create)

	app.Listen(":" + viper.GetString("app.port"))
}
