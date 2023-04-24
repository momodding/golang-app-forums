package controller

import (
	request "forum-app/model/request"
	"forum-app/model/response"
	"forum-app/service"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Register(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{userService: userService}
}

func (ctrl *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	var body request.UserRegistration
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	user, err := ctrl.userService.Register(body)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(
		response.NewSuccessResponse(user, "Register Success"),
	)
}
