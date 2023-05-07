package config

import (
	"forum-app/model/response"
	"forum-app/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

type Middleware struct {
	oauthService service.OauthService
}

func NewMiddleware(oauthService service.OauthService) *Middleware {
	return &Middleware{oauthService: oauthService}
}

func (mw *Middleware) OauthClientMiddleware(c *fiber.Ctx) error {
	clientId := c.Query("clientId")

	c.Locals("client", nil)
	if clientId != "" {
		log.Println("clientId = " + clientId)
		client, _ := mw.oauthService.GetClient(clientId)
		c.Locals("client", client)
	}

	return c.Next()
}

func (mw *Middleware) TokenValidatorMiddleware(c *fiber.Ctx) error {
	tokenHeader := c.Get("Authorization")
	tokenHeader = tokenHeader[1 : len(tokenHeader)-1]
	log.Printf("token: " + tokenHeader)
	accessToken := strings.TrimPrefix(tokenHeader, "Bearer ")
	err := mw.oauthService.Authenticate(accessToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewErrorResponse(fiber.StatusUnauthorized, err.Error(), "Authentication Failed!"))
	}

	return c.Next()
}
