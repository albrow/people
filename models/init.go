package models

import (
	"github.com/albrow/zoom"
	"log"
)

var pool *zoom.Pool

func init() {
	pool = zoom.NewPool(&zoom.PoolConfig{Database: 1})
	var err error
	People, err = pool.Register(&Person{})
	if err != nil {
		log.Fatal(err)
	}
}

func ClosePool() error {
	return pool.Close()
}
