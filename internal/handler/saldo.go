package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/gofiber/fiber/v2"
)

func GetSaldos(app *fiber.Ctx, db *database.Database) error {
	//Opcion 1
	/*
		resp, err := http.Get("https://pruebasmorpheus.com:18100/mrtire/medida/listarMedidas")
		if err != nil {
			return app.Status(500).SendString("Error al llamar a la API: " + err.Error())
		}
		defer resp.Body.Close()

		// 2. Leer el cuerpo de la respuesta como []byte
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return app.Status(500).SendString("Error al leer la respuesta: " + err.Error())
		}

		// 3. Parsear el JSON a un mapa (map[string]interface{})
		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return app.Status(500).SendString("Error al parsear JSON: " + err.Error())
		}
	*/

	//Opcion 2
	uri := "https://pruebasmorpheus.com:18100/mrtire/producto/listar"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal("Error al crear la peticion", err)
	}

	req.Header.Set("modelo_id", "0")
	req.Header.Set("medida_id", "0")
	req.Header.Set("marca_id", "4")

	// 3. Hacer la petición HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return app.Status(500).JSON(fiber.Map{"error": "Error al llamar a la API externa"})
	}
	defer resp.Body.Close()

	// 4. Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return app.Status(500).JSON(fiber.Map{"error": "Error al leer la respuesta"})
	}

	// 5. Parsear el JSON genérico
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return app.Status(500).JSON(fiber.Map{"error": "Error al parsear JSON"})
	}

	// 6. Devolver la respuesta al cliente
	return app.Status(200).JSON(data)

}
