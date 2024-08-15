package main

import (
	"90HM/pkg/server"
	"log"
)

func main() {
	srv := server.New()
	log.Println("Starting server on port 8080")
	log.Fatal(srv.ListenAndServe())
}