package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome page.</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>This is the contact page.</h1> "+
		"<p>email me at <a href='mailto:stantheman@gmail.com'>stantheman@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>FAQ</h1> "+
		"<p>"+
		"<strong>Q:</strong> Is there a free version?<br/>"+
		"<strong>A:</strong> Yes! We offer a free trial for 30 days on any paid plans<br/>"+
		"</p>"+
		"<p>"+
		"<strong>Q:</strong> What are your support hours?<br/>"+
		"<strong>A:</strong> We have support staff answering emails 24/7, though response times may be a bit slower on weekends.<br/>"+
		"</p>"+
		"<p>"+
		"<strong>Q:</strong> How do I contact support?<br/>"+
		"<strong>A:</strong> Email duh.<br/>"+
		"</p>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	})

	log.Println("server started")
	http.ListenAndServe(":3000", r)
}
