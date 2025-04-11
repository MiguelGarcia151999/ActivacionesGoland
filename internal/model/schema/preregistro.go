package schema

type Preregistro struct {
	Id               int    `json:"id"`
	Folio            string `json:"folio"`
	Digitos          string `json:"digitos"`
	Iccid            string `json:"iccid"`
	Fecha            string `json:"fecha"`
	Monto            int    `json:"monto"`
	Personalventa_id int    `json:"personalventa_id"`
	Carrier_id       int    `json:"carrier_id"`
	Tipochip_id      int    `json:"tipochip_id"`
	Lada_id          int    `json:"lada_id"`
}
