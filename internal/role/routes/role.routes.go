package routes

import (
	"bazar/internal/role/controllers"
	"bazar/internal/role/repositories"
	"bazar/internal/role/services"
	"bazar/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	db := config.GetDB()
	repo := repositories.NewRoleRepository(db)
	service := services.NewRoleService(repo)
	roleConroller := controllers.NewRoleController(service)

	role := app.Group("/role")

	role.Post("/create", roleConroller.Register)
}
