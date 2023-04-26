package controller

import (
	"forum-app/model/request"
	"forum-app/model/response"
	"forum-app/service"
	"github.com/gofiber/fiber/v2"
)

type OauthController interface {
	Authorize(ctx *fiber.Ctx) error
}

type OauthControllerImpl struct {
	oauthService service.OauthService
}

func NewOauthController(oauthService service.OauthService) *OauthControllerImpl {
	return &OauthControllerImpl{oauthService: oauthService}
}

func (ctrl *OauthControllerImpl) Authorize(ctx *fiber.Ctx) error {
	var body request.AuthorizationGrant
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	grantTypes := map[string]func(body request.AuthorizationGrant) response.AccessTokenResponse{
		"password": ctrl.oauthService.PasswordGrant,
	}

	handler, isHandlerExist := grantTypes[body.GrantType]
	if !isHandlerExist {
		return ctx.Status(fiber.StatusOK).JSON(
			response.NewErrorResponse(fiber.StatusBadRequest, fiber.Map{
				"grant_type": " Invalid grant type",
			}, "Authorization Failed"),
		)
	}
	result := handler(body)

	return ctx.Status(fiber.StatusOK).JSON(
		response.NewSuccessResponse(result, "Authorization Success"),
	)
}
