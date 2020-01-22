package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	Model "github.com/SGarcia710/comba-dashboard-server/Model/structs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	db := OpenDB()
	defer db.Close()

	// Routes
	router := mux.NewRouter()

	// Endpoints de software
	router.HandleFunc("/api/software", getSoftwares).Methods("GET")
	router.HandleFunc("/api/software", createSoftware).Methods("POST")
	// router.HandleFunc("/api/software/{id}", getSoftware).Methods("GET")
	// router.HandleFunc("/api/software/{id}", updateSoftware).Methods("PUT")
	// router.HandleFunc("/api/software/{id}", deleteSoftware).Methods("DELETE")

	// Endpoints de autores
	router.HandleFunc("/api/autor", getAuthors).Methods("GET")
	router.HandleFunc("/api/autor", createAuthor).Methods("POST")
	// router.HandleFunc("/api/autor/{id}", getAuthor).Methods("GET")
	// router.HandleFunc("/api/autor/{id}", updateAuthor).Methods("PUT")
	// router.HandleFunc("/api/autor/{id}", deleteAuthor).Methods("DELETE")

	fmt.Println("Listening at http://" + SERVER_DOMAIN + ":" + SERVER_PORT)
	http.ListenAndServe(":"+SERVER_PORT, router)
}

//	Author Handlers
func getAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authors []Model.Author

	result, err := db.Query("SELECT * from AUTOR")
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

func createAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENTRO AL HANDLER DE CREAR AUTHOR")

	stmt, err := db.Prepare("INSERT INTO autor(id_aut, ced_aut, nom_aut, ape_aut, sex_aut, cel_aut, email_aut, niv_acad_aut, ciud_aut, num_cvlac_aut, rol_aut, estado_aut, admin_aut) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("PREPARO LA QUERY DE INSERCIÓN")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(body)

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	cedula := keyVal["title"]
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

	_, err = stmt.Exec(nil, cedula, nombres, apellidos, sexo, celular, email, nivelAcademico, ciudad, numCvlac, rol, estado, admin)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New author was created")
}
