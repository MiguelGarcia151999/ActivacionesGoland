package handler

import (
	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/MiguelGarcia151999/ActivacionesGoland/internal/model/querys"
	"github.com/gofiber/fiber/v2"
)

func GetUsuarios(app *fiber.Ctx, db *database.Database) error {
	usuarios, err := querys.GetAllUsers(db)
	if err != nil {
		return app.Status(500).JSON(fiber.Map{
			"error": "Error al obtener usuarios",
		})
	}

	return app.Status(200).JSON(usuarios)
}
