package models

import (
	"github.com/albrow/zoom"
)

func Init() error {
	if err := zoom.Init(&zoom.Configuration{Database: 1}); err != nil {
		return err
	}
	var err error
	People, err = zoom.Register(&Person{})
	if err != nil {
		return err
	}
	return nil
}
