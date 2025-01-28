package main

import (
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	port := ":8080"

	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("src"))))

	handler := func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("src/index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Bladerunner", Director: "Ford Coppola"},
				{Title: "Blade Runnder", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)

	}

	http.HandleFunc("/", handler)

	http.ListenAndServe(port, nil)
}
