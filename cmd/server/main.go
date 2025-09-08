package main

import (
	chatbotRoutes "bazar/internal/chatbot/routes"
	roleRoutes "bazar/internal/role/routes"
	userRoutes "bazar/internal/users/routes"
	"bazar/pkg/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	config.ConnectDB()
	app := fiber.New()

	roleRoutes.SetUpRoutes(app)
	userRoutes.SetUpUsers(app)
	chatbotRoutes.SetUpChatbot(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
