package main

import (
	"log"
	"net/http"
	"tetris/cmd/routes"
)

func main() {
	// serve the css
	fs := http.FileServer(http.Dir("style"))
	http.Handle("GET /style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("GET /", routes.Root)
	http.HandleFunc("POST /tick", routes.Tick)
	http.HandleFunc("POST /restart", routes.Restart)

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
