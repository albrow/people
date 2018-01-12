package models

import (
	"log"

	"github.com/albrow/zoom"
)

var (
	People *zoom.Collection
	pool   *zoom.Pool
)

func init() {
	pool = zoom.NewPoolWithOptions(zoom.DefaultPoolOptions.WithDatabase(1))
	var err error
	People, err = pool.NewCollectionWithOptions(
		&Person{},
		zoom.DefaultCollectionOptions.WithIndex(true),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ClosePool() error {
	return pool.Close()
}
