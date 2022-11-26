package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to our test page</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Get in touch</h1>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(http.NotFound)
	http.ListenAndServe(":3333", r)
}
