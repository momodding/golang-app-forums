package main

import (
	"forum-app/config"
	"forum-app/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
	"log"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/categories", categoryController.FindAll)
	app.Post("/categories", categoryController.Create)

	app.Post("/users/register", userController.Register)

	oauthRoute := app.Group("/oauth")
	oauthRoute.Post("authorize", oauthController.Authorize)

	authRoute := app.Group("/auth", func(c *fiber.Ctx) error {
		clientId := c.Query("clientId")

		c.Locals("client", nil)
		if clientId != "" {
			log.Println("clientId = " + clientId)
			client, _ := oauthController.GetClient(clientId)
			c.Locals("client", client)
		}

		return c.Next()
	})
	authRoute.Get("login", authController.LoginView)

	app.Listen(":" + viper.GetString("app.port"))
}
