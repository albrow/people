package models

import (
	"github.com/albrow/zoom"
)

var People *zoom.ModelType

type Person struct {
	Name string
	Age  int
	zoom.RandomId
}
