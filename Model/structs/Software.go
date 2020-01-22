package Model

type Software struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Fecha       string `json:"fecha"`
}
