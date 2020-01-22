package main

import (
	"fmt"
	"net/http"

	Handlers "github.com/SGarcia710/comba-dashboard-server/Handlers"
	Utils "github.com/SGarcia710/comba-dashboard-server/Utils"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Software Endpoints
	router.HandleFunc("/api/software", Handlers.GetSoftwares).Methods("GET")
	router.HandleFunc("/api/software", Handlers.CreateSoftware).Methods("POST")
	// router.HandleFunc("/api/software/{id}", Handlers.GetSoftware).Methods("GET")
	// router.HandleFunc("/api/software/{id}", Handlers.UpdateSoftware).Methods("PUT")
	// router.HandleFunc("/api/software/{id}", Handlers.DeleteSoftware).Methods("DELETE")

	// Authors Endpoints
	router.HandleFunc("/api/autor", Handlers.GetAuthors).Methods("GET")
	router.HandleFunc("/api/autor", Handlers.CreateAuthor).Methods("POST")
	// router.HandleFunc("/api/autor/{id}", Handlers.GetAuthor).Methods("GET")
	// router.HandleFunc("/api/autor/{id}", Handlers.UpdateAuthor).Methods("PUT")
	// router.HandleFunc("/api/autor/{id}", Handlers.DeleteAuthor).Methods("DELETE")

	fmt.Println("Listening at http://" + Utils.SERVER_DOMAIN + ":" + Utils.SERVER_PORT)
	http.ListenAndServe(":"+Utils.SERVER_PORT, router)
}
