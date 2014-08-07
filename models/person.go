package models

import (
	"github.com/albrow/zoom"
	"github.com/mholt/binding"
)

type Person struct {
	Name string
	Age  int
	zoom.DefaultData
}

// Field mapping (pointer receiver is vital)
func (p *Person) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.Age: binding.Field{
			Form:     "age",
			Required: true,
		},
		&p.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
	}
}
