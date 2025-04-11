package querys

/* func GetLineasInfo(db *database.Database) ([]schema.Preregistro, error) {
	var lineas []schema.Preregistro
	err := db.Conn.Select(&lineas, "SELECT id, folio, digitos, iccid, IFNULL(fecha, 'S/R') AS fecha, monto, personalventa_id, carrier_id, tipochip_id, lada_id FROM preregistro LIMIT 10")
	return lineas, err
	// Retorna (datos, error) como es estándar en Go
} */

/* func GetLineasAll(db *database.Database) ([]map[string]interface{}, error) {
	rows, err := db.Conn.Queryx("SELECT * FROM preregistro LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		err := rows.MapScan(row)
		if err != nil {
			return nil, err
		}

		// Decodificar campos que están en Base64
		for key, value := range row {
			if strValue, ok := value.(string); ok {
				// Intenta decodificar si parece Base64
				if decoded, err := base64.StdEncoding.DecodeString(strValue); err == nil {
					// Verifica si el resultado decodificado es imprimible
					if isPrintable(decoded) {
						row[key] = string(decoded)
					}
				}
			}
		}

		results = append(results, row)
	}
	return results, nil
}

// isPrintable verifica si los bytes son caracteres imprimibles
func isPrintable(b []byte) bool {
	for _, c := range b {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
} */
