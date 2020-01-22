package Model

type Author struct {
	ID             int    `json:"id"`
	Cedula         string `json:"cedula"`
	Nombres        string `json:"nombres"`
	Apellidos      string `json:"apellidos"`
	Sexo           string `json:"sexo"`
	Celular        string `json:"celular"`
	Email          string `json:"email"`
	NivelAcademico string `json:"nivelAcademico"`
	Ciudad         string `json:"ciudad"`
	NumCvlac       string `json:"numCvlac"`
	Rol            string `json:"rol"`
	Estado         string `json:"estado"`
	Admin          int    `json:"admin"`
}
