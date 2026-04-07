package main

import (
	"log"
	"net/http"
	"os"
)

func serverAddress(port string) string {
	if port == "" {
		port = "8080"
	}

	return ":" + port
}

func run(getenv func(string) string, listenAndServe func(string, http.Handler) error) error {
	address := serverAddress(getenv("PORT"))
	log.Printf("starting server on %s", address)

	return listenAndServe(address, newMux())
}

func main() {
	if err := run(os.Getenv, http.ListenAndServe); err != nil {
		log.Fatal(err)
	}
}
