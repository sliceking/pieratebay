package main

import (
	"fmt"
	"log"
)
import "net/http"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome page.</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>This is the contact page.</h1> "+
		"<p>email me at <a href='mailto:stantheman@gmail.com'>stantheman@gmail.com</a></p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found")
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	log.Println("server started")
	http.ListenAndServe(":3000", nil)
}
