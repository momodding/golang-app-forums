package main

import (
	"forum-app/config"
	"forum-app/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

func main() {
	viewEngine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		Views:        viewEngine,
	})
	app.Use(recover.New())

	config.LoadConfig()

	categoryController := InitializeCategoryController()
	userController := InitializeUserController()
	oauthController := InitializeOauthController()
	authController := InitializeAuthController()

	app.Use(func(c *fiber.Ctx) error {
		clientId := c.Get("clientId")
		client, _ := oauthController.GetClient(clientId)
		c.Locals("client", client)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/categories", categoryController.FindAll)
	app.Post("/categories", categoryController.Create)

	app.Post("/users/register", userController.Register)

	oauthRoute := app.Group("/oauth")
	oauthRoute.Post("authorize", oauthController.Authorize)

	authRoute := app.Group("/auth")
	authRoute.Get("login", authController.LoginView)

	app.Listen(":" + viper.GetString("app.port"))
}
