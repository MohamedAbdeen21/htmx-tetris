package views

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed *.html
var viewsFS embed.FS

var tmpl *template.Template

func init() {
	var err error

	tmpl, err = template.ParseFS(viewsFS, "*.html")
	if err != nil {
		log.Fatal(http.StatusInternalServerError, err.Error())
	}
}

func Render(w http.ResponseWriter, block string, data any) {
	if err := tmpl.ExecuteTemplate(w, block, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
