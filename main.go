package main

import (
	"log"
)

func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	store.Init()

	server := NewAPIServer(":8000", store)
	server.Run()
}
