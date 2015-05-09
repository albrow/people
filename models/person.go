package models

import (
	"fmt"
	"github.com/albrow/zoom"
)

var People *zoom.ModelType

type Person struct {
	Name string
	Age  int
	zoom.DefaultData
}

func (p Person) MarshalJSON() ([]byte, error) {
	data := fmt.Sprintf(`{"Id": "%s", "Name": "%s", "Age": %d}`, p.Id(), p.Name, p.Age)
	return []byte(data), nil
}
