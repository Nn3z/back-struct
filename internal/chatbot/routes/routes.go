package routes

import (
	"fmt"

	"bazar/internal/chatbot/controllers"
	"bazar/internal/chatbot/service"

	"github.com/gofiber/fiber/v2"
)

func SetUpChatbot(app *fiber.App) {
	// 1️⃣ Crear el ChatbotService leyendo tu JSON
	chatbotService, err := service.NewChatbotService("db/cv.json")
	if err != nil {
		panic(fmt.Sprintf("No se pudo inicializar el ChatbotService: %v", err))
	}

	// 2️⃣ Crear el controller pasando el service
	controller := controllers.NewChatbotController(chatbotService)

	// 3️⃣ Configurar la ruta POST /chatbot
	chatbot := app.Group("/chatbot")
	chatbot.Post("/", controller.Ask)
}
