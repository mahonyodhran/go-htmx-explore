package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting app...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!")
		io.WriteString(w, r.Method)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8008", nil))
}
