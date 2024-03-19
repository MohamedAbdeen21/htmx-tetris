package main

import (
	"htmx/cmd/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", routes.Root)
	http.HandleFunc("POST /update", routes.Update)

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
