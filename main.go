package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := ":" + port
	log.Printf("starting server on %s", address)

	if err := http.ListenAndServe(address, newMux()); err != nil {
		log.Fatal(err)
	}
}
