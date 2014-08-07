package controllers

import (
	"github.com/albrow/go-data-parser"
	"github.com/albrow/learning/peeps-negroni/models"
	"github.com/albrow/zoom"
	"github.com/unrolled/render"
	"net/http"
)

type Persons struct{}

func (c Persons) Create(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	userData, err := data.Parse(req)
	if err != nil {
		panic(err)
	}

	// validations
	val := userData.Validator()
	val.Require("name")
	val.Require("age")
	if val.HasErrors() {
		r.JSON(res, 400, newJSONValidationError(val))
		return
	}

	// save to database
	p := &models.Person{
		Name: userData.Get("name"),
		Age:  userData.GetInt("age"),
	}
	if err := zoom.Save(p); err != nil {
		panic(err)
	}

	r.JSON(res, 200, p)
}
