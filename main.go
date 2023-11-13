package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting app...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8008", nil))
}
