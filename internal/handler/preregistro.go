package handler

import (
	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/gofiber/fiber/v2"
)

func GetLineas(app *fiber.Ctx, db *database.Database) error {
	/* lineas, err := querys.GetLineas(db)
	if err != nil {
		return app.Status(500).JSON(fiber.Map{
			"error": "Error al obtener lineas",
		})
	}

	return app.Status(200).JSON(lineas) */

	//lineas, err := querys.GetLineasInfo(db)
	/* if err != nil {
		return app.Status(500).JSON(fiber.Map{
			"error": "Error al obtener lineas",
		})
	} */

	return app.Status(200).JSON("OK")
}
