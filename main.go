package main

import (
	"fmt"
	"log"
)
import "net/http"

func SomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to this awesome site")
}
func main() {
	http.HandleFunc("/", SomeHandler)
	log.Println("server started")
	http.ListenAndServe(":3000", nil)
}
