package main

import (
	"embed"
	"log"
	"net/http"
	"tetris/cmd/routes"
)

//go:embed style/*
var content embed.FS

func main() {
	// serve the css
	var fs = http.FileServer(http.FS(content))
	http.Handle("GET /style/", fs)

	http.HandleFunc("GET /", routes.Root)
	http.HandleFunc("POST /tick", routes.Tick)
	http.HandleFunc("POST /restart", routes.Restart)

	log.Print("Starting server on port 8080")
	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
