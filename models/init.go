package models

import (
	"github.com/albrow/zoom"
)

func Init() error {
	zoom.Init(&zoom.Configuration{Database: 1})

	var err error
	People, err = zoom.Register(&Person{})
	if err != nil {
		return err
	}
	return nil
}
