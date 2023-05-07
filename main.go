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

	middleware := InitializeMiddleware()
	categoryController := InitializeCategoryController()
	userController := InitializeUserController()
	oauthController := InitializeOauthController()
	authController := InitializeAuthController()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/categories", middleware.TokenValidatorMiddleware, categoryController.FindAll)
	app.Post("/categories", middleware.TokenValidatorMiddleware, categoryController.Create)

	app.Post("/users/register", userController.Register)

	oauthRoute := app.Group("/oauth")
	oauthRoute.Post("authorize", oauthController.Authorize)

	authRoute := app.Group("/auth", middleware.OauthClientMiddleware)
	authRoute.Get("login", authController.LoginView)

	err := app.Listen(":" + viper.GetString("app.port"))
	if err != nil {
		log.Fatal("Error running app: " + err.Error())
	}
}
