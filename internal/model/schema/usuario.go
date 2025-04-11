package schema

type Usuario struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Nick   string `json:"nick"`
}
