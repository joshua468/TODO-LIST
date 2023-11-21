package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

var tmpl *template.Template

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn GO", Done: false},
			{Item: "Like TEMI", Done: false},
		},
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("./static", http.StripPrefix("./static/", fs))
	mux.HandleFunc("/todo", todo)
	log.Fatal(http.ListenAndServe(":9091", mux))

}
