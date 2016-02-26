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
	pool = zoom.NewPool(&zoom.PoolOptions{Database: 1})
	var err error
	People, err = pool.NewCollection(&Person{}, &zoom.CollectionOptions{
		Index: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ClosePool() error {
	return pool.Close()
}
