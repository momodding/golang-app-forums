package controller

import (
	"forum-app/helper"
	request "forum-app/model/request"
	"forum-app/model/response"
	"forum-app/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Register(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	userService service.UserService
	Validate    *validator.Validate
}

func NewUserController(userService service.UserService, Validate *validator.Validate) *UserControllerImpl {
	return &UserControllerImpl{userService: userService, Validate: Validate}
}

func (ctrl *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	var body request.UserRegistration
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	ctrl.Validate.RegisterValidation("validateUsername", ctrl.userService.ValidateUsername)
	err := ctrl.Validate.Struct(body)
	helper.PanicIfError(err)

	user, err := ctrl.userService.Register(body)
	helper.PanicIfError(err)

	return ctx.Status(fiber.StatusOK).JSON(
		response.NewSuccessResponse(user, "Register Success"),
	)
}
