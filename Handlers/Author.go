package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	Model "github.com/SGarcia710/comba-dashboard-server/Model/structs"
	Utils "github.com/SGarcia710/comba-dashboard-server/Utils"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	db := Utils.OpenDB()
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var authors []Model.Author

	result, err := db.Query("SELECT * from autor")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	for result.Next() {
		var autor Model.Author
		err := result.Scan(&autor.ID, &autor.Cedula, &autor.Nombres, &autor.Apellidos, &autor.Sexo, &autor.Celular, &autor.Email, &autor.NivelAcademico, &autor.Ciudad, &autor.NumCvlac, &autor.Rol, &autor.Estado, &autor.Admin)
		if err != nil {
			panic(err.Error())
		}
		authors = append(authors, autor)
	}
	json.NewEncoder(w).Encode(authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	db := Utils.OpenDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO autor(id_aut, ced_aut, nom_aut, ape_aut, sex_aut, cel_aut, email_aut, niv_acad_aut, ciud_aut, num_cvlac_aut, rol_aut, estado_aut, admin_aut) VALUES(null, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	cedula := keyVal["cedula"]
	nombres := keyVal["nombres"]
	apellidos := keyVal["apellidos"]
	sexo := keyVal["sexo"]
	celular := keyVal["celular"]
	email := keyVal["email"]
	nivelAcademico := keyVal["nivelAcademico"]
	ciudad := keyVal["ciudad"]
	numCvlac := keyVal["numCvlac"]
	rol := keyVal["rol"]
	estado := keyVal["estado"]
	admin := keyVal["admin"]

	_, err = stmt.Exec(cedula, nombres, apellidos, sexo, celular, email, nivelAcademico, ciudad, numCvlac, rol, estado, admin)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New author was created")
}
