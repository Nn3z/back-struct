package main

import (
	roleRoutes "bazar/internal/role/routes"
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

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
