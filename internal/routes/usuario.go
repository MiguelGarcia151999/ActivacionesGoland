package routes

import (
	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/MiguelGarcia151999/ActivacionesGoland/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func UsuarioRoutes(app *fiber.App, db *database.Database) {
	api := app.Group("/usuario")

	api.Get("/listar", func(c *fiber.Ctx) error {
		return handler.GetUsuarios(c, db)
	})
}
