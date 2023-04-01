package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	categoryController := InitializeCategoryController()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/categories", categoryController.FindAll)
	app.Post("/categories", categoryController.Create)

	app.Listen(":3000")
}
