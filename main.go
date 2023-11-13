package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	log.Println("Starting app...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "ET", Director: "Stephen Spielberg"},
				{Title: "21 Jump Street", Director: "Phil Lord"},
				{Title: "The Prestige", Director: "Christopher Nolan"},
			},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8008", nil))
}
