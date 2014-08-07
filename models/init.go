package models

import (
	"github.com/albrow/zoom"
)

func Init() {
	zoom.Init(&zoom.Configuration{Database: 1})
	zoom.Register(&Person{})
}
