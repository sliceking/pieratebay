package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sliceking/pieratebay/views"
	"log"
	"net/http"
	"path/filepath"
)

func executeTemplate(w http.ResponseWriter, tPath string) {
	t, err := views.Parse(tPath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tPath)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	})

	log.Println("server started")
	http.ListenAndServe(":3000", r)
}
