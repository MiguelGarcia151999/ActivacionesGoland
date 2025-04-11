package routes

import (
	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *database.Database) {
	UsuarioRoutes(app, db)
	LineasRoutes(app, db)
	SaldoRoutes(app, db)
}
