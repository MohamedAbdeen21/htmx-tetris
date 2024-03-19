package routes

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.New("")
var count = 0

func init() {
	var err error

	tmpl, err = template.ParseGlob("views/*.html")

	if err != nil {
		log.Fatal(http.StatusInternalServerError, err.Error())
	}
}

func render(w http.ResponseWriter, block string, data any) {
	if err := tmpl.ExecuteTemplate(w, block, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
