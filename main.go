package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/MiguelGarcia151999/ActivacionesGoland/internal/routes"
	"github.com/MiguelGarcia151999/ActivacionesGoland/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Error al conectar a la DB: %v", err)
	}

	routes.SetupRoutes(app, db)
	middleware.SetDefaultRoutes(app)

	go func() {
		// Mensaje de intento de inicio (útil para debugging)
		fmt.Println("⚡ Iniciando servidor HTTPS en el puerto 19070...")
		fmt.Println("\033[32m🟢 Servidor HTTPS iniciado correctamente\033[0m")

		err := app.ListenTLS(
			":19070",
			"/etc/letsencrypt/live/pruebasmorpheus.com/fullchain.pem",
			"/etc/letsencrypt/live/pruebasmorpheus.com/privkey.pem",
		)
		if err != nil {
			log.Fatalf("\033[31m🚨 Error al iniciar HTTPS: %v\033[0m", err) // Fatal detiene el programa
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("📡 Recibida señal de apagado...")
	if err := app.Shutdown(); err != nil {
		log.Printf("⚠️ Error al apagar el servidor: %v", err)
	}
	log.Println("🔴 Servidor apagado correctamente")
}
