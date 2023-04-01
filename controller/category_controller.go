package controller

import (
	"forum-app/entity"
	"forum-app/service"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	FindAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{service: service}
}

func (ctrl *CategoryControllerImpl) FindAll(c *fiber.Ctx) error {
	categories := ctrl.service.FindAll()
	return c.Status(fiber.StatusOK).JSON(categories)
}

func (ctrl *CategoryControllerImpl) Create(c *fiber.Ctx) error {
	var request entity.Category
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	category := ctrl.service.Save(request)
	return c.Status(fiber.StatusOK).JSON(category)
}
