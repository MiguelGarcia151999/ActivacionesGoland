package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/MiguelGarcia151999/ActivacionesGoland/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Database representa la conexión a MySQL
type Database struct {
	Conn *sqlx.DB
}

func ConnectDB() (*Database, error) {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("DB_PORT debe ser un número válido") // Retorna (nil, error)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"))

	conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("❌ Error conectando a la base de datos: %w", err)
	}

	log.Println("✅ Conexión a MySQL establecida correctamente")
	return &Database{Conn: conn}, nil
}
