package config

import (
	"forum-app/service"
	"github.com/gofiber/fiber/v2"
	"log"
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
