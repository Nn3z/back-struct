package routes

import (
	"bazar/internal/users/controllers"
	"bazar/internal/users/repositories"
	"bazar/internal/users/services"
	"bazar/pkg/config"
	"bazar/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpUsers(app *fiber.App) {
	db := config.GetDB()
	repo := repositories.NewUserRepository(db)
	services := services.NewUserService(repo)
	userController := controllers.NewUserController(services)

	user := app.Group("/auth")

	user.Post("/signup", userController.SignUp)
	user.Post("/login", userController.Login)
	user.Get("/profile", middleware.JWTmiddleware, userController.Profile)
}
