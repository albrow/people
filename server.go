package main

import (
	"github.com/albrow/peeps-negroni/controllers"
	"github.com/albrow/peeps-negroni/models"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/albrow/negroni-json-recovery"
)

func main() {
	models.Init()

	router := mux.NewRouter()
	persons := controllers.Persons{}
	router.HandleFunc("/persons", persons.Create).Methods("POST")
	router.HandleFunc("/persons/{id}", persons.Show).Methods("GET")
	router.HandleFunc("/persons/{id}", persons.Update).Methods("PUT", "PATCH")
	router.HandleFunc("/persons", persons.Index).Methods("GET")
	router.HandleFunc("/persons/{id}", persons.Delete).Methods("DELETE")

	n := negroni.New(negroni.NewLogger())
	n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router)

	n.Run(":3000")
}
