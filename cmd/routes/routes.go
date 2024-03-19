package routes

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	// TODO: Middleware to stop longest match
	if r.URL.Path != "/" {
		http.Error(w, "No route", http.StatusInternalServerError)
		return
	}

	render(w, "index", count)
}

func Update(w http.ResponseWriter, r *http.Request) {
	count++
	render(w, "counter", count)
}
