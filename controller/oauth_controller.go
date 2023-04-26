package controller

import (
	"forum-app/helper"
	"forum-app/model/request"
	"forum-app/model/response"
	"forum-app/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OauthController interface {
	Authorize(ctx *fiber.Ctx) error
}

type OauthControllerImpl struct {
	oauthService service.OauthService
	Validate     *validator.Validate
}

func NewOauthController(oauthService service.OauthService, Validate *validator.Validate) *OauthControllerImpl {
	return &OauthControllerImpl{oauthService: oauthService, Validate: Validate}
}

func (ctrl *OauthControllerImpl) Authorize(ctx *fiber.Ctx) error {
	var body request.AuthorizationGrant
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	ctrl.Validate.RegisterValidation("validateGrantType", ctrl.oauthService.ValidateGrantType)
	err := ctrl.Validate.Struct(body)
	helper.PanicIfError(err)

	grantTypes := map[string]func(body request.AuthorizationGrant) response.AccessTokenResponse{
		"password": ctrl.oauthService.PasswordGrant,
	}

	handler, isHandlerExist := grantTypes[body.GrantType]
	if !isHandlerExist {
		return ctx.Status(fiber.StatusOK).JSON(
			response.NewErrorResponse(fiber.StatusBadRequest, fiber.Map{
				"grant_type": " invalid grant type",
			}, "Authorization Failed"),
		)
	}
	result := handler(body)

	return ctx.Status(fiber.StatusOK).JSON(
		response.NewSuccessResponse(result, "Authorization Success"),
	)
}
