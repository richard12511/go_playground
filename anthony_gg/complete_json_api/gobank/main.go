package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalf(err.Error())
	}

	server := NewAPIServer(":8000", store)
	server.Run()
}