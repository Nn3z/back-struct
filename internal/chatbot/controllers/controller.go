package controllers

import (
	dtos "bazar/internal/chatbot/DTOs"
	"bazar/internal/chatbot/service"

	"github.com/gofiber/fiber/v2"
)

type ChatbotController struct {
	service *service.ChatbotService
}

func NewChatbotController(service *service.ChatbotService) *ChatbotController {
	return &ChatbotController{service: service}
}

func (c *ChatbotController) Ask(ctx *fiber.Ctx) error {
	var body dtos.AskRequestDTO

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	answer, err := c.service.Ask(body.Question)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dtos.AskResponseDTO{
		Answer: answer,
	}

	return ctx.JSON(response)
}
