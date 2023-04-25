package exception

import (
	"forum-app/model/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if isError, resp := validationError(c, err); isError {
		return resp
	}

	return internalServerError(c, err)
}

func validationError(ctx *fiber.Ctx, err error) (bool, error) {
	exception, isError := err.(validator.ValidationErrors)
	if isError {
		errorOut := make(map[string]string, len(exception))
		for _, fieldErr := range exception {
			errorOut[fieldErr.Field()] = ParseTags(fieldErr)
		}
		code := fiber.StatusBadRequest
		resp := ctx.Status(code).JSON(response.NewErrorResponse(code, errorOut, "error validation"))
		return true, resp
	}

	return false, nil
}

func internalServerError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	resp := ctx.Status(code).JSON(response.NewErrorResponse(code, err.Error(), err.Error()))

	return resp
}
