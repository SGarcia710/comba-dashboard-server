package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	Model "github.com/SGarcia710/comba-dashboard-server/Model/structs"
	Utils "github.com/SGarcia710/comba-dashboard-server/Utils"
)

func GetSoftwares(w http.ResponseWriter, r *http.Request) {
	db := Utils.OpenDB()
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var softwares []Model.Software

	result, err := db.Query("SELECT * from software")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var software Model.Software
		err := result.Scan(&software.ID, &software.Nombre, &software.Descripcion, &software.Fecha)
		if err != nil {
			panic(err.Error())
		}
		softwares = append(softwares, software)
	}
	json.NewEncoder(w).Encode(softwares)
}

func CreateSoftware(w http.ResponseWriter, r *http.Request) {
	db := Utils.OpenDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO software(id_soft, nom_soft, desc_soft, fec_soft) VALUES(null, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nombre := keyVal["nombre"]
	descripcion := keyVal["descripcion"]
	fecha := keyVal["fecha"]

	_, err = stmt.Exec(nombre, descripcion, fecha)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New software was created")
}
