package routes

import (
	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/MiguelGarcia151999/ActivacionesGoland/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SaldoRoutes(app *fiber.App, db *database.Database) {
	api := app.Group("/saldo")

	api.Get("/SaldoActual", func(c *fiber.Ctx) error {
		return handler.GetSaldos(c, db)
	})
}
