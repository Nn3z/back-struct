package controllers

import (
	"bazar/internal/role/model"
	roleservice "bazar/internal/role/services"

	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	service *roleservice.RoleService
}

func NewRoleController(service *roleservice.RoleService) *RoleController {
	return &RoleController{service: service}
}

func (rc *RoleController) Register(c *fiber.Ctx) error {
	role := new(model.Role)
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"error":   err.Error(),
		})
	}
	if err := rc.service.CreateRole(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(role)
}
