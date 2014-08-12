package controllers

import (
	"github.com/albrow/go-data-parser"
	"github.com/albrow/learning/peeps-negroni/models"
	"github.com/albrow/zoom"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

type Persons struct{}

func (c Persons) Create(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	personData, err := data.Parse(req)
	if err != nil {
		panic(err)
	}

	// validations
	val := personData.Validator()
	val.Require("name")
	val.Require("age")
	if val.HasErrors() {
		r.JSON(res, 400, newJSONValidationError(val))
		return
	}

	// save to database
	p := &models.Person{
		Name: personData.Get("name"),
		Age:  personData.GetInt("age"),
	}
	if err := zoom.Save(p); err != nil {
		panic(err)
	}

	// render response
	dataMap := map[string]interface{}{"Person": p}
	r.JSON(res, 200, newJSONData(dataMap))
}

func (c Persons) Show(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	vars := mux.Vars(req)
	id := vars["id"]

	// find in the database
	p := &models.Person{}
	if err := zoom.ScanById(id, p); err != nil {
		panic(err)
	}
	defer p.Unlock()

	// render response
	dataMap := map[string]interface{}{"Person": p}
	r.JSON(res, 200, newJSONData(dataMap))
}

func (c Persons) Update(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	vars := mux.Vars(req)
	id := vars["id"]

	// parse data
	personData, err := data.Parse(req)
	if err != nil {
		panic(err)
	}

	// find in the database
	p := &models.Person{}
	if err := zoom.ScanById(id, p); err != nil {
		panic(err)
	}
	defer p.Unlock()

	// update model
	if personData.KeyExists("name") {
		p.Name = personData.Get("name")
	}
	if personData.KeyExists("age") {
		p.Age = personData.GetInt("age")
	}
	if err := zoom.Save(p); err != nil {
		panic(err)
	}

	// render response
	dataMap := map[string]interface{}{"Person": p}
	r.JSON(res, 200, newJSONData(dataMap))
}

func (c Persons) Index(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})

	// find all persons in the database
	var persons []*models.Person
	if err := zoom.NewQuery("Person").Scan(&persons); err != nil {
		panic(err)
	}
	for _, p := range persons {
		p.Unlock()
	}

	// render response
	dataMap := map[string]interface{}{"Persons": persons}
	r.JSON(res, 200, newJSONData(dataMap))
}

func (c Persons) Delete(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	vars := mux.Vars(req)
	id := vars["id"]

	// delete from database
	if err := zoom.DeleteById("Person", id); err != nil {
		panic(err)
	}

	// render response
	r.JSON(res, 200, newJSONOk())
}
