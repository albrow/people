package models

import (
	"github.com/albrow/zoom"
)

type Person struct {
	Name string
	Age  int
	zoom.DefaultData
}
