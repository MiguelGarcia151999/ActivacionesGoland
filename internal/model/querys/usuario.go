package querys

import (
	"github.com/MiguelGarcia151999/ActivacionesGoland/database"
	"github.com/MiguelGarcia151999/ActivacionesGoland/internal/model/schema"
)

// GetAllUsers obtiene todos los usuarios
func GetAllUsers(db *database.Database) ([]schema.Usuario, error) {
	var usuarios []schema.Usuario
	err := db.Conn.Select(&usuarios, "SELECT id, nombre, nick FROM usuario")
	return usuarios, err
	// Retorna (datos, error) como es estándar en Go
}
