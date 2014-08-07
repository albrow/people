package controllers

import (
	"github.com/albrow/learning/peeps-negroni/models"
	"github.com/albrow/zoom"
	"github.com/mholt/binding"
	"github.com/unrolled/render"
	"net/http"
)

type Persons struct{}

func (c Persons) Create(res http.ResponseWriter, req *http.Request) {
	p := new(models.Person)
	if binding.Bind(req, p).Handle(res) {
		return
	}

	if err := zoom.Save(p); err != nil {
		panic(err)
	}

	r := render.New(render.Options{})
	r.JSON(res, 200, p)
}
