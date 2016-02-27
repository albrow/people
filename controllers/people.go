package controllers

import (
	"fmt"
	"net/http"

	"github.com/albrow/zoom"

	"github.com/albrow/forms"
	"github.com/albrow/people/models"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type People struct{}

func (c People) Create(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	personData, err := forms.Parse(req)
	if err != nil {
		panic(err)
	}

	// validations
	val := personData.Validator()
	val.Require("Name")
	val.Require("Age")
	if val.HasErrors() {
		r.JSON(res, 422, val.ErrorMap())
		return
	}

	// save to database
	p := &models.Person{
		Name: personData.Get("Name"),
		Age:  personData.GetInt("Age"),
	}
	if err := models.People.Save(p); err != nil {
		panic(err)
	}

	// render response
	r.JSON(res, 200, p)
}

func (c People) Show(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	vars := mux.Vars(req)
	id := vars["id"]

	// find in the database
	p := &models.Person{}
	if err := models.People.Find(id, p); err != nil {
		if _, ok := err.(zoom.ModelNotFoundError); ok {
			r.JSON(res, 404, map[string]interface{}{
				"error": fmt.Sprintf("Could not find Person with id = %s", id),
			})
			return
		}
		panic(err)
	}

	// render response
	r.JSON(res, 200, p)
}

func (c People) Update(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	vars := mux.Vars(req)
	id := vars["id"]

	// parse data
	personData, err := forms.Parse(req)
	if err != nil {
		panic(err)
	}

	// find in the database
	p := &models.Person{}
	if err := models.People.Find(id, p); err != nil {
		panic(err)
	}

	// update model
	if personData.KeyExists("Name") {
		p.Name = personData.Get("Name")
	}
	if personData.KeyExists("Age") {
		p.Age = personData.GetInt("Age")
	}
	if err := models.People.Save(p); err != nil {
		panic(err)
	}

	// render response
	r.JSON(res, 200, p)
}

func (c People) Index(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})

	// find all people in the database
	people := []*models.Person{}
	if err := models.People.NewQuery().Run(&people); err != nil {
		panic(err)
	}

	// render response
	r.JSON(res, 200, people)
}

func (c People) Delete(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	vars := mux.Vars(req)
	id := vars["id"]

	// delete from database
	if _, err := models.People.Delete(id); err != nil {
		panic(err)
	}

	// render response
	r.JSON(res, 200, struct{}{})
}
