package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sliceking/pieratebay/controllers"
	"github.com/sliceking/pieratebay/templates"
	"github.com/sliceking/pieratebay/views"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"faq.gohtml", "tailwind.gohtml"))))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	//r.Post("/users", controllers.)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	})

	log.Println("server started")
	http.ListenAndServe(":3000", r)
}
