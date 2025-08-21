package controllers

import (
	dtos "bazar/internal/users/DTOs"
	userService "bazar/internal/users/services"
	"bazar/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	service *userService.UserService
}

func NewUserController(service *userService.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) SignUp(c *fiber.Ctx) error {
	user := new(dtos.UserSignUpDTO)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"error":   err.Error(),
		})
	}
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = string(hashedpassword)

	if user.Username == "" {
		nameNospaces := utils.RemoveSpaces(user.Name)
		randomCode, err := utils.GenerateRandomID(6)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "could not generate username",
			})
		}
		user.Username = nameNospaces + randomCode
	}
	id, err := utils.GenerateRandomID(16)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not generate user id",
		})
	}

	user.ID = id

	roleID := "1"
	user.RoleID = roleID
	if err := uc.service.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
