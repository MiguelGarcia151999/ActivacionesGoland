package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configura todas las rutas de la API
func SetDefaultRoutes(app *fiber.App) {
	// Rutas de usuarios
	// Ruta para manejar 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"error":   "Ruta no encontrada",
			"message": "La ruta solicitada no existe en este servidor",
		})
	})
}
