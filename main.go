package main

import (
	"log"

	"github.com/albrow/negroni-json-recovery"
	"github.com/albrow/people/controllers"
	"github.com/albrow/people/models"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	defer func() {
		if err := models.ClosePool(); err != nil {
			log.Fatal(err)
		}
	}()

	router := mux.NewRouter()
	people := controllers.People{}
	router.HandleFunc("/people", people.Index).Methods("GET")
	router.HandleFunc("/people", people.Create).Methods("POST")
	router.HandleFunc("/people/{id}", people.Show).Methods("GET")
	router.HandleFunc("/people/{id}", people.Update).Methods("PUT", "PATCH")
	router.HandleFunc("/people/{id}", people.Delete).Methods("DELETE")

	n := negroni.New(negroni.NewLogger())
	n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router)

	n.Run(":3000")
}
