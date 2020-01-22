package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	SERVER_DOMAIN = "localhost"        // Server domain
	SERVER_PORT   = "8000"             // Server port
	DB_HOST       = "localhost"        // DB Server Hostname/IP
	DB_PORT       = "3306"             // DB Port for MySQL
	DB_USER       = "comba"            /*"root"*/ // DB username
	DB_PASS       = "combapw"          /*""*/     // DB Password
	DB_NAME       = "combadashboarddb" // Database name
)

type Autor struct {
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

type Software struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Fecha       string `json:"fecha"`
}

type Desarrollo struct {
	ID         int `json:"id"`
	IDAutor    int `json:"idAutor"`
	IDSoftware int `json:"idSoftware"`
}

var db *sql.DB
var err error

func OpenDB() *sql.DB {
	fmt.Println("Estoy intentando conectarme a la base de datos...")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))
	// db, err = sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conexión a la base de datos establecida.")
	}
	return db
}

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
	var authors []Autor

	result, err := db.Query("SELECT * from AUTOR")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	for result.Next() {
		var autor Autor
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
		fmt.Println(err)
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

// Sofware Handlers
func getSoftwares(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var softwares []Software

	result, err := db.Query("SELECT * from software")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var software Software
		err := result.Scan(&software.ID, &software.Nombre, &software.Descripcion, &software.Fecha)
		if err != nil {
			panic(err.Error())
		}
		softwares = append(softwares, software)
	}
	json.NewEncoder(w).Encode(softwares)
} // Funciona

func createSoftware(w http.ResponseWriter, r *http.Request) {
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
} // Funciona
