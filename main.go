package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	addr := ":8080"

	http.HandleFunc("/", handler_one)
	http.HandleFunc("/add-film/", handler_two)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler_two(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	htmlStr := fmt.Sprintf("<li>%s -- %s</li>", title, director)

	tmpl, _ := template.New("t").Parse(htmlStr)

	tmpl.Execute(w, nil)
}

func handler_one(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("index.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "F. Ford Coppola"},
			{Title: "Blade Runner", Director: "Riddley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}
