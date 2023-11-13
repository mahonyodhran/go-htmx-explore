package main

import (
	"fmt"
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

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'> %s - %s </li>", title, director)
		template.New("t").Parse(htmlStr)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8008", nil))
}
