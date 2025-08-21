package routes

import (
	"bazar/internal/users/controllers"
	"bazar/internal/users/repositories"
	"bazar/internal/users/services"
	"bazar/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func SetUpUsers(app *fiber.App) {
	db := config.GetDB()
	repo := repositories.NewUserRepository(db)
	services := services.NewUserService(repo)
	userController := controllers.NewUserController(services)

	user := app.Group("/auth")

	user.Post("/signup", userController.SignUp)

}
